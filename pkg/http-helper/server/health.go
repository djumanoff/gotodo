package server

import (
	"encoding/json"
	"net/http"
)

type Healther interface {
	Health() error
}

func healthHandler(healthers ...Healther) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var errs []error
		for _, h := range healthers {
			if err := h.Health(); err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			d, err := json.Marshal(errs)
			if err != nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			_, err = w.Write(d)
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	}
}
