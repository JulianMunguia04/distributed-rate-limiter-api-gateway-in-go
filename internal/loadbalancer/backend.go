package loadbalancer

import (
	"net/url"
	"sync"
	"sync/atomic"

	"gateway/internal/circuitbreaker"
)

type Backend struct {
	URL *url.URL

	Alive bool
	mu    sync.RWMutex

	Connections int64
	CB          *circuitbreaker.CircuitBreaker
}

func (b *Backend) SetAlive(alive bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Alive = alive
}

func (b *Backend) IsAlive() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Alive
}

func (b *Backend) IncConnections() {
	atomic.AddInt64(&b.Connections, 1)
}

func (b *Backend) DecConnections() {
	atomic.AddInt64(&b.Connections, -1)
}

func (b *Backend) GetConnections() int64 {
	return atomic.LoadInt64(&b.Connections)
}
