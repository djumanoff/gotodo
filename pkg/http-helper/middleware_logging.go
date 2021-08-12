package http_helper

import (
	"go.uber.org/zap"
	"net/http"
)

func Log(log zap.Logger, handler Handler) Handler {
	return func(r *http.Request) Response {
		log.Debug(r.Method)
		resp := handler(r)
		log.Debug(r.URL.String())
		return resp
	}
}
