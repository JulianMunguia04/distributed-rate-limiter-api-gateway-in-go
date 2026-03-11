package middleware

import (
	"net/http"
	"os"
)

func APIKeyAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		apiKey := os.Getenv("X_API_KEY")
		key := r.Header.Get("X_API_KEY")

		if key != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
