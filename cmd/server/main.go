package main

import (
	"github.com/djumanoff/gotodo/pkg/config"
	"github.com/djumanoff/gotodo/pkg/cqrses"
	hh "github.com/djumanoff/gotodo/pkg/http-helper"
	"github.com/djumanoff/gotodo/pkg/logger"
	"github.com/djumanoff/gotodo/pkg/todo"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	flags = []cli.Flag{
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
			EnvVars: []string{"ADDRESS"},
		},
	}

	commands = []*cli.Command{
		{
			Name:    "server",
			Aliases: []string{"run"},
			Usage:   "run http server",
			Action:  run,
			Flags:   flags,
		},
	}
)

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

type Config struct {
	Addr string `envconfig:"addr" mapstructure:"addr" default:":8080"`
}

func run(c *cli.Context) error {
	cfg := &Config{}
	configPath := c.String("config")
	if configPath != "" {
		_ = config.LoadFromFile("", cfg, configPath)
	}
	lg := logger.New()
	hhCfg := hh.Config{Addr: cfg.Addr, Logger: lg}
	router := hh.NewRouter(hhCfg)

	cmder := cqrses.NewCommandHandler(todo.NewService(todo.NewMockRepo()))
	errSys := hh.NewErrorSystem("TODO")

	fac := todo.NewHttpHandlerFactory(cmder, errSys)
	mw := hh.HttpMiddlewareFactory{}

	router.Mux.Get("/todos", mw.JSON(fac.GetTodos()))

	return hh.Listen(hhCfg, router)
}
