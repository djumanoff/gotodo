package http_helper

import (
	"net/http"
)

type Handler func(r *http.Request) Response

type Middleware func(handler Handler) Handler
