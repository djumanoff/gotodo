package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func WithLogging(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var fields []zap.Field
			for name, values := range r.Header {
				fields = append(fields, zap.Field{
					Key:    name,
					String: values[0],
					Type:   zapcore.StringType,
				})
			}
			rec := httptest.NewRecorder()
			next.ServeHTTP(rec, r)

			dump, _ := httputil.DumpResponse(rec.Result(), false)

			// we copy the captured response headers to our new response
			for k, v := range rec.Header() {
				w.Header()[k] = v
			}

			// grab the captured response body
			data := rec.Body.Bytes()
			w.WriteHeader(rec.Result().StatusCode)
			w.Write(data)

			fields = append(fields, zap.Field{
				Key:    "method",
				String: r.Method,
				Type:   zapcore.StringType,
			}, zap.Field{
				Key:    "url",
				String: r.URL.String(),
				Type:   zapcore.StringType,
			})
			log.Debug(string(dump)+" <::> "+string(data), fields...)
		}
		return http.HandlerFunc(fn)
	}
}
