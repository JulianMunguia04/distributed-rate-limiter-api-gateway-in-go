package main

import (
	"fmt"
	"log"
	"net/http"

	"gateway/internal/config"
	"gateway/internal/loadbalancer"
	"gateway/internal/middleware"
	"gateway/internal/proxy"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env:", err)
	} else {
		log.Println(".env loaded successfully")
	}

	// Load YAML config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	handlers := make(map[string]http.Handler)

	// Loop through services
	for _, svc := range cfg.Services {

		// Create load balancer per service
		lb := loadbalancer.NewLeastConnections(svc.Backends)

		// Reverse proxy handler
		handler := proxy.ProxyRequest(lb)

		var finalHandler http.Handler

		// Middleware chain (dynamic based on auth_required)
		switch svc.AuthRequired {
		case true:
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
				middleware.APIKeyAuth,
				middleware.RateLimit,
			)
		default:
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
				middleware.RateLimit,
			)
		}

		handlers[svc.Name] = finalHandler
	}

	// Register routes dynamically
	for name, handler := range handlers {
		route := "/" + name + "/"
		log.Println("Registering route:", route)
		http.Handle(route, handler)
	}

	// Start gateway on config port
	addr := ":" + fmt.Sprint(cfg.Port)
	log.Println("Gateway running on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
