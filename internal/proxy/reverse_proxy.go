package proxy

import (
	"net/http"
	"net/http/httputil"

	"gateway/internal/loadbalancer"
)

func ProxyRequest(lb loadbalancer.LoadBalancer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		backend := lb.NextBackend()

		if backend == nil {
			http.Error(w, "No healthy backends", http.StatusServiceUnavailable)
			return
		}

		backend.IncConnections()
		defer backend.DecConnections()

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)
		proxy.ServeHTTP(w, r)
	})
}
