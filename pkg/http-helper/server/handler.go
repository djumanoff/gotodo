package server

import (
	"github.com/djumanoff/gotodo/pkg/http-helper"
	"net/http"
)

type Handler func(r *http.Request) http_helper.Response

type Middleware func(handler Handler) Handler

type OutputMiddleware func(handler Handler) http.HandlerFunc
