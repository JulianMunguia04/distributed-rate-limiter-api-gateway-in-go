package middleware

import "net/http"

func RateLimit(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// later you will call your token bucket here

		next.ServeHTTP(w, r)
	})
}
