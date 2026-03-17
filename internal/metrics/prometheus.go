package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gateway_requests_total",
		Help: "Total number of HTTP requests processed by the gateway",
	},
	[]string{"path", "method"},
)

var RequestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "gateway_request_duration_seconds",
		Help:    "Latency of HTTP requests",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"path"},
)

var ErrorsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gateway_errors_total",
		Help: "Total number of error responses",
	},
	[]string{"path", "status"},
)

func init() {
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(ErrorsTotal)
}
