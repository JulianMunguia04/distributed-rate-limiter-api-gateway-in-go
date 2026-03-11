package middleware

import (
	"net/http"
	"os"
)

var apiKey = os.Getenv("API_KEY")

func APIKeyAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.Header.Get("X-API-KEY")

		if key != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
