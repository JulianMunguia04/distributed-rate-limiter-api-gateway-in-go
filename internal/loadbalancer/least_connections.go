package loadbalancer

import (
	"log"
	"math"
	"net/url"

	"gateway/internal/circuitbreaker"
)

type LeastConnections struct {
	backends []*Backend
}

func NewLeastConnections(urls []string) *LeastConnections {
	var backends []*Backend

	for _, rawURL := range urls {
		u, err := url.Parse(rawURL)
		if err != nil {
			log.Println("Invalid backend URL:", rawURL)
			continue
		}

		backends = append(backends, &Backend{
			URL:   u,
			Alive: true,
			CB:    circuitbreaker.NewCircuitBreaker(rawURL),
		})
	}

	return &LeastConnections{
		backends: backends,
	}
}

func (lb *LeastConnections) GetBackends() []*Backend {
	return lb.backends
}

func (lb *LeastConnections) NextBackend() *Backend {
	var selected *Backend
	min := int64(math.MaxInt64)

	for _, b := range lb.backends {

		// Skip dead backends
		if !b.IsAlive() {
			continue
		}

		// Skip circuit breaker blocked
		if b.CB != nil && !b.CB.CanRequest() {
			continue
		}

		conn := b.GetConnections()

		if conn < min {
			min = conn
			selected = b
		}
	}

	return selected
}
