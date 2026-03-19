package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"sync/atomic"

	"gateway/internal/loadbalancer"
)

func ProxyRequest(lb loadbalancer.LoadBalancer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		backend := lb.NextBackend()

		log.Println("Forwarding to:", backend.URL)

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)

		// success response
		proxy.ModifyResponse = func(resp *http.Response) error {
			atomic.AddInt64(&backend.Connections, -1)

			if resp.StatusCode >= 500 {
				backend.CB.OnFailure()
			} else {
				backend.CB.OnSuccess()
			}

			return nil
		}

		// failure case
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			atomic.AddInt64(&backend.Connections, -1)

			log.Println("Error:", err)

			backend.CB.OnFailure()

			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		}

		proxy.ServeHTTP(w, r)
	}
}
