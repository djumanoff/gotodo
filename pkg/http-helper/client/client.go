package client

import (
	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/hystrix"
	"time"
)

type Config struct {
	HttpTimeout int `envconfig:"http_timeout" default:"300"` // 300ms

	// Hystrix Circuit Breaker strategy
	HystrixTimeout   int `envconfig:"hystrix_timeout" default:"1100"` // 1.1s
	MaxConcurrent    int `envconfig:"max_concurrent" default:"100"`
	ErrorThreshold   int `envconfig:"error_threshold" default:"25"`
	Sleep            int `envconfig:"sleep" default:"10"` // sleep window
	RequestThreshold int `envconfig:"request_threshold" default:"10"`

	// Exponential Backoff Retry strategy
	InitialTimeout int     `envconfig:"initial_timeout" defualt:"2"` // 2ms
	MaxTimeout     int     `envconfig:"max_timeout" default:"1000"`  // 1s
	ExponentFactor float64 `envconfig:"exponent_factor" default:"2"` // Multiplier
	MaxJitter      int     `envconfig:"max_jitter" default:"2"`      // 2ms (it should be more than 1ms)
	RetryCount     int     `envconfig:"retry_count" default:"4"`
}

func NewRetryer(cfg Config) heimdall.Retriable {
	initTimeout := time.Duration(cfg.InitialTimeout) * time.Millisecond
	maxTimeout := time.Duration(cfg.MaxTimeout) * time.Millisecond
	maxJitter := time.Duration(cfg.MaxJitter) * time.Millisecond

	return heimdall.NewRetrier(heimdall.NewExponentialBackoff(
		initTimeout,
		maxTimeout,
		cfg.ExponentFactor,
		maxJitter,
	))
}

// TODO Create client from a factory with different type of clients
func NewClient(service string, cfg Config, client heimdall.Doer) *hystrix.Client {
	timeout := time.Duration(cfg.HttpTimeout) * time.Millisecond
	if client != nil {
		timeout = 0 * time.Millisecond
	}

	hystrixTimeout := time.Duration(cfg.HystrixTimeout) * time.Millisecond
	clt := hystrix.NewClient(
		hystrix.WithHTTPTimeout(timeout),
		hystrix.WithCommandName(service),
		hystrix.WithHystrixTimeout(hystrixTimeout),
		hystrix.WithMaxConcurrentRequests(cfg.MaxConcurrent),
		hystrix.WithErrorPercentThreshold(cfg.ErrorThreshold),
		hystrix.WithSleepWindow(cfg.Sleep),
		hystrix.WithRequestVolumeThreshold(cfg.RequestThreshold),
		hystrix.WithHTTPClient(client),
		hystrix.WithRetrier(NewRetryer(cfg)),
		hystrix.WithRetryCount(cfg.RetryCount),
	)
	return clt
}
