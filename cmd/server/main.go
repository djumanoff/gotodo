package main

import (
	"github.com/djumanoff/gotodo/internal/todo"
	config "github.com/l00p8/cfg"
	"github.com/l00p8/cqrses"
	"github.com/l00p8/l00p8"
	logger "github.com/l00p8/log"
	hh "github.com/l00p8/xserver"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

var (
	// commands and flags of the cli
	commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"run"},
			Usage:   "run http server",
			Action:  run,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Usage:   "Load configuration from .env file",
					Value:   "",
					EnvVar: "CONFIG",
				},
				&cli.StringFlag{
					Name:    "address",
					Usage:   "run http server on specified address",
					Value:   "",
					EnvVar: "ADDR",
				},
			},
		},
	}
)

// initialize application
func main() {
	app := &cli.App{}
	app.Name = "Todo Server"
	app.Usage = "Http server that handles all Todo app use cases."
	app.Commands = commands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// Config struct for server command
type Config struct {
	Addr           string `envconfig:"addr" mapstructure:"addr" default:":8080"`
	RateLimit      int64  `envconfig:"rate_limit" mapstructure:"rate_limit" default:"1"`
	DBFile         string `envconfig:"db_file" mapstructure:"db_file" default:"db.sqlite"`
	MigrationsFile string `envconfig:"migrations_file" mapstructure:"migrations_file" default:""`
	LogLevel       string `envconfig:"log_level" mapstructure:"log_level" default:"debug"`
	System         string `envconfig:"system" mapstructure:"system" default:"TODO"`
	Hostname       string `envconfig:"hostname" mapstructure:"hostname" default:"localhost"`
}

func (cfg *Config) load(c *cli.Context) {
	// load config from file if config file provided
	configPath := c.String("config")
	if configPath != "" {
		_ = config.LoadFromFile("", cfg, configPath)
	}
	// load config from command line arguments
	if cfg.Addr == "" {
		cfg.Addr = c.String("address")
	}
}

// run func runs http server, returns error if configuration is invalid for some reason
func run(c *cli.Context) error {
	cfg := &Config{}

	// init config
	cfg.load(c)

	lg, err := logger.NewLogger(cfg.LogLevel, zap.String("system", cfg.System), zap.String("hostname", cfg.Hostname))
	must(err)
	logFac := logger.NewFactory(lg)

	// init config for http server
	hhCfg := hh.Config{
		GracefulTimeout: 3 * time.Second,
		ShutdownTimeout: 3 * time.Second,
		Addr:            cfg.Addr,
		RateLimit:       cfg.RateLimit,
		Logger:          logFac,
	}
	router := l00p8.NewHandlerRouter(
		hh.NewRouterWithTracing(hh.NewRouter(hhCfg)),
		l00p8.JSON,
	)

	repo, err := todo.NewSqliteRepository(todo.SqliteConfig{
		DbName:         "todos",
		FilePath:       cfg.DBFile,
		MigrationsFile: cfg.MigrationsFile,
	})
	must(err)

	cmder := cqrses.NewCommandHandlerWithPublisher(
		cqrses.NewKafkaPublisher(cqrses.KafkaConfig{}),
		cqrses.NewCommandHandler(todo.NewService(repo)),
	)

	// init error system
	errSys := l00p8.NewErrorSystem("TODO")
	fac := todo.NewHttpHandlerFactory(cmder, errSys)

	// init routes
	router.Get("/todos", fac.GetTodos())
	router.Post("/todos", fac.CreateTodo())
	router.Get("/todos/{id}", fac.GetTodo("id"))

	// init health checks
	router.Healthers(repo)

	// start http server with cleanup function
	// to close db connections, files, queues etc.
	return hh.Listen(hhCfg, router, func() {
		lg.Info("cleanup func called")
		time.Sleep(1 * time.Second)
		lg.Info("cleanup finished")
	})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
