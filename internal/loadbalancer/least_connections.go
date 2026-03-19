package loadbalancer

import (
	"net/url"
	"sync/atomic"

	"gateway/internal/circuitbreaker"
)

type LeastConnections struct {
	backends []*Backend
}

func NewLeastConnections(addresses []string) *LeastConnections {

	var backends []*Backend

	for _, addr := range addresses {
		parsed, err := url.Parse(addr)
		if err != nil {
			continue
		}

		backends = append(backends, &Backend{
			URL:         parsed,
			Connections: 0,
			CB:          circuitbreaker.NewCircuitBreaker(parsed.String()),
		})
	}

	return &LeastConnections{
		backends: backends,
	}
}

func (lb *LeastConnections) NextBackend() *Backend {

	var selected *Backend

	for _, backend := range lb.backends {

		// 🔴 skip unhealthy backends
		if backend.CB != nil && !backend.CB.CanRequest() {
			continue
		}

		if selected == nil ||
			atomic.LoadInt64(&backend.Connections) <
				atomic.LoadInt64(&selected.Connections) {

			selected = backend
		}
	}

	// fallback if all are open
	if selected == nil {
		selected = lb.backends[0]
	}

	atomic.AddInt64(&selected.Connections, 1)

	return selected
}
