package server

import (
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func WithLogging(log *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			//var fields []zap.Field
			//for name, values := range r.Header {
			//	fields = append(fields, zap.String("req_"+name, values[0]))
			//}
			rec := httptest.NewRecorder()
			next.ServeHTTP(rec, r)

			dumpResp, _ := httputil.DumpResponse(rec.Result(), false)
			dumpReq, _ := httputil.DumpRequest(r, true)

			// we copy the captured response headers to our new response
			for k, v := range rec.Header() {
				w.Header()[k] = v
			}

			// grab the captured response body
			data := rec.Body.Bytes()
			w.WriteHeader(rec.Result().StatusCode)
			_, _ = w.Write(data)

			//fields = append(fields,
			//	zap.String("method", r.Method),
			//	zap.String("url", r.URL.String()),
			//	zap.Int("response_size", n),
			//)
			log.Debug("", zap.String("request", string(dumpReq)), zap.String("response", string(dumpResp)))
			//log.Debug(string(dumpResp)+" <::> "+string(data), fields...)
		}
		return http.HandlerFunc(fn)
	}
}
