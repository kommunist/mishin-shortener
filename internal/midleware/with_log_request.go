package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Мидлварь логирования запроса
func WithLogRequest(h http.Handler) http.Handler {
	logfn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData: &responseData{
				status: 0,
				size:   0,
			},
		}
		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		slog.Info(
			"Request",
			"REQUEST_METHOD", r.Method,
			"URL", r.URL,
			"Duration", duration,
			"RESPONSE_STATUS", lw.responseData.status,
			"SIZE", lw.responseData.size,
		)

	}

	return http.HandlerFunc(logfn)

}
