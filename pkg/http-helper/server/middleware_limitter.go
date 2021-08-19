package server

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"net/http"
)

func cancelCtx(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{}"))
	if err == nil {
		return
	}
	w.WriteHeader(http.StatusTooManyRequests)
}

func rateLimitter(lmt *limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-r.Context().Done():
				cancelCtx(w, r)
			default:
				next.ServeHTTP(w, r)
			}
		}
		return tollbooth.LimitHandler(lmt, http.HandlerFunc(fn))
	}
}
