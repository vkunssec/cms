package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
		startTime := time.Now()

		lrw := NewLoggingResponseWriter(w)

		next.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode

		duration := time.Since(startTime).Round(time.Microsecond)
		log.Printf("%d.%d status_code=%d %s=%s response_latency=%s", r.ProtoMajor, r.ProtoMinor, statusCode, r.Method, r.RequestURI, duration)
	})
}
