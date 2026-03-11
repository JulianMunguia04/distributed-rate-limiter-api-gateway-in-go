package loadbalancer

import (
	"log"
	"net/url"
	"sync"
)

type RoundRobin struct {
	backends []*Backend
	mu       sync.Mutex
	current  int
}

func NewRoundRobin(addresses []string) *RoundRobin {

	var backends []*Backend

	for _, addr := range addresses {

		parsed, err := url.Parse(addr)
		if err != nil {
			log.Fatal(err)
		}

		backends = append(backends, &Backend{
			URL: parsed,
		})
	}

	return &RoundRobin{
		backends: backends,
	}
}

func (lb *RoundRobin) NextBackend() *Backend {

	lb.mu.Lock()
	defer lb.mu.Unlock()

	backend := lb.backends[lb.current]

	lb.current = (lb.current + 1) % len(lb.backends)

	return backend
}
