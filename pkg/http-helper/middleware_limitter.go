package http_helper

import (
	"github.com/didip/tollbooth"
	"net/http"
)

func cancelCtx(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("{}"))
	if err == nil {
		return
	}
	w.WriteHeader(http.StatusTooManyRequests)
}

func (fac *HttpMiddlewareFactory) RateLimit(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			cancelCtx(w, r)
		default:
			next.ServeHTTP(w, r)
		}
	}
	return tollbooth.LimitHandler(fac.RateLimitter, http.HandlerFunc(fn))
}
