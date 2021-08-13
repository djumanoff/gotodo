package http_helper

import "github.com/didip/tollbooth/limiter"

type HttpMiddlewareFactory struct {
	RateLimitter *limiter.Limiter
}
