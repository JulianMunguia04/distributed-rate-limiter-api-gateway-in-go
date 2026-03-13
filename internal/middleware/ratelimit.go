package middleware

import (
	"net"
	"net/http"

	"gateway/internal/ratelimiter"
)

func RateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Unable to determine IP", http.StatusInternalServerError)
			return
		}

		bucket := ratelimiter.GetBucket(ip)

		if !bucket.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
