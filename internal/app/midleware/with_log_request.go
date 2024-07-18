package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func WithLogRequest(h http.Handler) http.Handler {
	logfn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		duration := time.Since(start)

		slog.Info("Request", "METHOD", r.Method, "URL", r.URL, "Duration", duration)

	}

	return http.HandlerFunc(logfn)

}
