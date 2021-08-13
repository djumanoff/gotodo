package http_helper

import (
	"encoding/json"
	"net/http"
)

func (fac *HttpMiddlewareFactory) JSON(handler Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		resp := handler(r)
		data, err := json.Marshal(resp.Response())
		if err != nil {
			_, err = w.Write([]byte("{}"))
			if err == nil {
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err == nil {
			return
		}
		w.WriteHeader(resp.StatusCode())
	}
	return fn
}
