package http_helper

import (
	"context"
	"encoding/json"
	"net/http"
)

func (fac *HttpMiddlewareFactory) JSON(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		resp := handler(ctx, r)
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
}
