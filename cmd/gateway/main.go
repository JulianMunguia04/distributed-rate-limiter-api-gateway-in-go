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

	http.Handle("/metrics", promhttp.Handler())

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env:", err)
	}

	// Load YAML config
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	handlers := make(map[string]http.Handler)

	// Loop through services
	for _, svc := range cfg.Services {

		// Create load balancer
		lb := loadbalancer.NewLeastConnections(svc.Backends)

		// Reverse proxy handler
		handler := proxy.ProxyRequest(lb)

		var finalHandler http.Handler

		// Apply middleware
		switch svc.Name {
		case "app1":
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
				middleware.RateLimit,
			)
		case "app2":
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
				middleware.APIKeyAuth,
				middleware.RateLimit,
			)
		default:
			// Fallback if a service is added but no middleware defined yet
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
			)
			log.Printf("Warning: No specific middleware defined for service %s", svc.Name)
		}

		handlers[svc.Name] = finalHandler
	}

	// Register routes dynamically
	for name, handler := range handlers {
		route := "/" + name + "/"
		log.Println("Registering route:", route)
		http.Handle(route, handler)
	}

	addr := ":" + fmt.Sprint(cfg.Port)

	log.Println("Gateway running on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
