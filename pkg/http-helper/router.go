package http_helper

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"time"
)

type Resource interface {
	Routes() *Router
}

type Router struct {
	Mux    chi.Router
	Config Config
}

//func NewRouter() *Router {
//	return &Router{Mux: chi.NewRouter()}
//}

func NewApiRouter(apiVersion string) *Router {
	r := &Router{Mux: chi.NewRouter()}
	return r
}

func (r *Router) AddResource(pattern string, resource Resource) {
	r.Mux.Mount(pattern, resource.Routes().Mux)
}

func (r *Router) Prefix(prefix string, router *Router) {
	r.Mux.Mount(prefix, router.Mux)
}

func (r *Router) Healthers(healthers ...Healther) {
	r.Mux.Get("/_health", healthHandler(healthers...))
}

func NewRouter(cfg Config) *Router {
	r := &Router{Mux: chi.NewRouter(), Config: cfg}

	timeout := r.Config.Timeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	r.Mux.Use(chiMiddleware.RequestID)
	//r.Mux.Use(middleware.Logger(logger.L()))
	r.Mux.Use(chiMiddleware.StripSlashes)
	r.Mux.Use(chiMiddleware.Recoverer)
	r.Mux.Use(chiMiddleware.Timeout(timeout))

	return r
}
