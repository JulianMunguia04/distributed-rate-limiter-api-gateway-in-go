package loadbalancer

import (
	"net/url"
)

type Backend struct {
	URL         *url.URL
	Connections int64
}

type LoadBalancer interface {
	NextBackend() *Backend
}
