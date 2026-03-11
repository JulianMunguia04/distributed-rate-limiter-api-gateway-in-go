package proxy

import (
	"net/http"
	"net/http/httputil"

	"gateway/internal/loadbalancer"
)

func ProxyRequest(lb loadbalancer.LoadBalancer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		backend := lb.NextBackend()

		proxy := httputil.NewSingleHostReverseProxy(backend.URL)

		proxy.ServeHTTP(w, r)
	}
}
