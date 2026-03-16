package middleware

import (
	"net"
	"net/http"
	"time"

	"gateway/internal/ratelimiter"
)

func RateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "IP error", http.StatusInternalServerError)
			return
		}

		allowed, err := ratelimiter.AllowRequest(ip, 100, time.Minute)
		if err != nil {
			http.Error(w, "Rate limiter error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
