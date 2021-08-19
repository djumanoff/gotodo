package main

import (
	"github.com/djumanoff/gotodo/pkg/config"
	"github.com/djumanoff/gotodo/pkg/cqrses"
	hh "github.com/djumanoff/gotodo/pkg/http-helper"
	"github.com/djumanoff/gotodo/pkg/http-helper/server"
	"github.com/djumanoff/gotodo/pkg/logger"
	"github.com/djumanoff/gotodo/pkg/todo"
	"github.com/djumanoff/gotodo/pkg/todo/sqlite"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var (
	// commands and flags of the cli
	commands = []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"run"},
			Usage:   "run http server",
			Action:  run,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "config",
					Aliases: []string{"c"},
					Usage:   "Load configuration from .env file",
					Value:   "",
					EnvVars: []string{"CONFIG", "CFG"},
				},
				&cli.StringFlag{
					Name:    "address",
					Aliases: []string{"a"},
					Usage:   "run http server on specified address",
					Value:   "",
					EnvVars: []string{"ADDR"},
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

	lg := logger.New()
	lg.SetLevel(cfg.LogLevel)
	mw := server.HttpMiddlewareFactory{}

	// init config for http server
	hhCfg := server.Config{
		GracefulTimeout: 3 * time.Second,
		ShutdownTimeout: 3 * time.Second,
		Addr:            cfg.Addr,
		RateLimit:       cfg.RateLimit,
		Logger:          lg,
	}
	router := server.NewRouterWithOutput(hhCfg, mw.JSON)

	repo, err := sqlite.NewRepository(sqlite.Config{
		DbName:         "todos",
		FilePath:       cfg.DBFile,
		MigrationsFile: cfg.MigrationsFile,
	})
	must(err)

	cmder := cqrses.NewCommandHandler(todo.NewService(repo))

	// init error system
	errSys := hh.NewErrorSystem("TODO")

	fac := todo.NewHttpHandlerFactory(cmder, errSys)

	// init routes
	router.Get("/todos", fac.GetTodos())
	router.Post("/todos", fac.CreateTodo())
	router.Get("/todos/{id}", fac.GetTodo("id"))

	// init health checks
	router.Healthers(repo)

	// start http server with cleanup function
	// to close db connections, files, queues etc.
	return server.Listen(hhCfg, router, func() {
		lg.Logger.Info("cleanup func called")
		time.Sleep(3 * time.Second)
		lg.Logger.Info("cleanup finished")
	})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
