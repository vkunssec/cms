package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponse struct {
	http.ResponseWriter
	statusCode int
}

func responseWriter(w http.ResponseWriter) *loggingResponse {
	return &loggingResponse{w, http.StatusOK}
}

func (lrw *loggingResponse) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	startTime := time.Now()

	lrw := responseWriter(w)

	statusCode := lrw.statusCode

	duration := time.Since(startTime).Round(time.Microsecond)
	log.Printf("%d.%d status_code=%d %s=%s response_latency=%s", r.ProtoMajor, r.ProtoMinor, statusCode, r.Method, r.RequestURI, duration)

	next(lrw, r)
}
