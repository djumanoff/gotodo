package http_helper

import (
	"context"
	"net/http"
)

type Handler func(ctx context.Context, r *http.Request) Response

type Middleware func(handler Handler) Handler
