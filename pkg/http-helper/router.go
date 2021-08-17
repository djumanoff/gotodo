package http_helper

import (
	"github.com/didip/tollbooth"
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
	mw     OutputMiddleware
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

func (r *Router) Get(prefix string, fn Handler) {
	r.Mux.Get(prefix, r.mw(fn))
}

func (r *Router) Post(prefix string, fn Handler) {
	r.Mux.Post(prefix, r.mw(fn))
}

func (r *Router) Put(prefix string, fn Handler) {
	r.Mux.Put(prefix, r.mw(fn))
}

func (r *Router) Delete(prefix string, fn Handler) {
	r.Mux.Delete(prefix, r.mw(fn))
}

func NewRouter(cfg Config) *Router {
	r := &Router{Mux: chi.NewRouter(), Config: cfg}
	lmt := tollbooth.NewLimiter(float64(cfg.RateLimit), nil)

	timeout := r.Config.Timeout
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	r.Mux.Use(chiMiddleware.RequestID)
	//r.Mux.Use(middleware.Logger(logger.L()))
	r.Mux.Use(chiMiddleware.StripSlashes)
	r.Mux.Use(chiMiddleware.Recoverer)
	r.Mux.Use(chiMiddleware.Timeout(timeout))
	r.Mux.Use(prometheusMiddleware)
	r.Mux.Use(rateLimitter(lmt))

	return r
}

func NewRouterWithOutput(cfg Config, mw OutputMiddleware) *Router {
	r := NewRouter(cfg)
	r.mw = mw
	return r
}
