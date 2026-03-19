package loadbalancer

import (
	"net/url"

	"gateway/internal/circuitbreaker"
)

type Backend struct {
	URL         *url.URL
	Connections int64
	CB          *circuitbreaker.CircuitBreaker
}

type LoadBalancer interface {
	NextBackend() *Backend
}
