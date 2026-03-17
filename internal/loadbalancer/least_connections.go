package loadbalancer

import (
	"net/url"
	"sync/atomic"
)

type LeastConnections struct {
	backends []*Backend
}

func NewLeastConnections(addresses []string) *LeastConnections {

	var backends []*Backend

	for _, addr := range addresses {
		parsed, _ := url.Parse(addr)

		backends = append(backends, &Backend{
			URL:         parsed,
			Connections: 0,
		})
	}

	return &LeastConnections{
		backends: backends,
	}
}

func (lb *LeastConnections) NextBackend() *Backend {

	var selected *Backend

	for _, backend := range lb.backends {

		if selected == nil ||
			atomic.LoadInt64(&backend.Connections) <
				atomic.LoadInt64(&selected.Connections) {

			selected = backend
		}
	}

	atomic.AddInt64(&selected.Connections, 1)

	return selected
}
