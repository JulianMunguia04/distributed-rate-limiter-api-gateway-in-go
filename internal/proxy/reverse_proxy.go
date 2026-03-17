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

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)

		// decrement after success
		proxy.ModifyResponse = func(resp *http.Response) error {
			atomic.AddInt64(&backend.Connections, -1)
			return nil
		}

		// decrement on error
		proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			atomic.AddInt64(&backend.Connections, -1)
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		}

		log.Println("Forwarding request to:", backend.URL)

		proxy.ServeHTTP(w, r)
	}
}
