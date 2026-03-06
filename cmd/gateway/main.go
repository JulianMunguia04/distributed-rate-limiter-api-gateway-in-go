package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL *url.URL
}

type LoadBalancer struct {
	backends []*Backend
	mu       sync.Mutex
	current  int
}

func (lb *LoadBalancer) NextBackend() *Backend {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	backend := lb.backends[lb.current]
	lb.current = (lb.current + 1) % len(lb.backends)

	return backend
}

func newLoadBalancer(addresses []string) *LoadBalancer {
	var backends []*Backend

	for _, addr := range addresses {
		parsed, err := url.Parse(addr)
		if err != nil {
			log.Fatal(err)
		}
		backends = append(backends, &Backend{URL: parsed})
	}

	return &LoadBalancer{
		backends: backends,
	}
}

func main() {

	app1LB := newLoadBalancer([]string{
		"http://localhost:5000",
		"http://localhost:5001",
	})

	app2LB := newLoadBalancer([]string{
		"http://localhost:3000",
		"http://localhost:3001",
	})

	http.HandleFunc("/app1/", func(w http.ResponseWriter, r *http.Request) {
		backend := app1LB.NextBackend()
		proxy := httputil.NewSingleHostReverseProxy(backend.URL)
		proxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/app2/", func(w http.ResponseWriter, r *http.Request) {
		backend := app2LB.NextBackend()
		proxy := httputil.NewSingleHostReverseProxy(backend.URL)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//go run main.go