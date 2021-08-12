package http_helper

import (
	"context"
	"github.com/djumanoff/gotodo/pkg/config"
	"github.com/djumanoff/gotodo/pkg/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/valve"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config describes server configuration
type Config struct {
	Addr            string `envconfig:"addr" mapstructure:"addr" default:":8080"`
	ShutdownTimeout int    `envconfig:"shutdown_timeout" mapstructure:"shutdown_timeout" default:"20"`
	GracefulTimeout int    `envconfig:"graceful_timeout" mapstructure:"graceful_timeout" default:"21"`
	HealthUri       string `envconfig:"health_uri" mapstructure:"health_uri" default:"/_health"`
	ApiVersion      string `envconfig:"api_version" mapstructure:"api_version" default:"v1"`
	Timeout         int    `envconfig:"timeout" mapstructure:"timeout" default:"20"`
}

func init() {
	// initialize defaults for the Viper
	config.SetDefault(constant.ApiVersionCtxKey, constant.ApiVersionDefault)
	config.SetDefault(constant.HealthUriKey, "/_health")
	config.SetDefault(constant.TimeoutKey, 20)
}

// Listen starts a http server on specified address and defines gateway routes
// Server implements a graceful shutdown pattern for better handling of rolling k8s updates
func Listen(cfg Config, router *Router) {
	valv := valve.New()
	baseCtx := valv.Context()

	srv := http.Server{
		Addr:    cfg.Addr,
		Handler: chi.ServerBaseContext(baseCtx, router.Mux),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.S().Info("Shutting down a http server...\n")

			shutdown := time.Duration(cfg.ShutdownTimeout) * time.Second

			// first valv
			if err := valv.Shutdown(shutdown); err != nil {
				logger.S().Errorf("Error shutting down a Valve: %v", err)
				return
			}

			// create a context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), shutdown)
			defer cancel()

			// start http server shutdown
			if err := srv.Shutdown(ctx); err != nil {
				logger.S().Errorf("Error shutting down a http server: %v", err)
				return
			}

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(time.Duration(cfg.GracefulTimeout) * time.Second):
				logger.S().Info("Not all connections are done")
			case <-ctx.Done():

			}
		}
	}()

	logger.S().Infof("Starting a new server on address: %s", cfg.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.S().Errorf("A server listener error: %v", err)
	}

	logger.S().Info("Server is down")
}
