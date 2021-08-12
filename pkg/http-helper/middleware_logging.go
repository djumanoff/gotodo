package http_helper

import (
	"context"
	"go.uber.org/zap"
	"net/http"
)

func (fac *HttpMiddlewareFactory) Log(log zap.Logger, handler Handler) Handler {
	return func(ctx context.Context, r *http.Request) Response {
		log.Debug(r.Method)
		resp := handler(ctx, r)
		log.Debug(r.URL.String())
		return resp
	}
}
