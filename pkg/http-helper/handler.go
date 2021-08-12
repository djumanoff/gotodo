package http_helper

import (
	"context"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type Handler func(ctx context.Context, r *http.Request) Response

type Middleware func(handler Handler) Handler
