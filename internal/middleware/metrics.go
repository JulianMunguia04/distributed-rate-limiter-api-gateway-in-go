package middleware

import (
	"net/http"
	"time"

	"gateway/internal/metrics"
)

func Metrics(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		metrics.RequestsTotal.WithLabelValues(r.URL.Path, r.Method).Inc()

		next.ServeHTTP(w, r)

		duration := time.Since(start).Seconds()

		metrics.RequestDuration.WithLabelValues(r.URL.Path).Observe(duration)
	})
}
