package loadbalancer

import "net/url"

type Backend struct {
	URL *url.URL
}

type LoadBalancer interface {
	NextBackend() *Backend
}
