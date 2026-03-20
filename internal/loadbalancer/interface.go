package loadbalancer

type LoadBalancer interface {
	NextBackend() *Backend
	GetBackends() []*Backend
}
