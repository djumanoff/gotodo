package http_helper

import (
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

type Handler func(r *http.Request) Response

type Middleware func(handler Handler) Handler
