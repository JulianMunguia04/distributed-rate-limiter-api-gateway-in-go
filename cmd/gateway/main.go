package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gateway/internal/config"
	"gateway/internal/loadbalancer"
	"gateway/internal/middleware"
	"gateway/internal/proxy"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// Metrics endpoint
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	// Load .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env:", err)
	} else {
		log.Println(".env loaded successfully")
	}

	// Load YAML config
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Dynamic services
	for _, svc := range cfg.Services {

		lb := loadbalancer.NewLeastConnections(svc.Backends)
		handler := proxy.ProxyRequest(lb)

		var finalHandler http.Handler

		if svc.AuthRequired {
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
				middleware.APIKeyAuth,
				middleware.RateLimit,
			)
		} else {
			finalHandler = middleware.Chain(
				handler,
				middleware.Metrics,
				middleware.Logging,
				middleware.Recovery,
				middleware.RateLimit,
			)
		}

		route := "/" + svc.Name + "/"
		log.Println("Registering route:", route)
		mux.Handle(route, finalHandler)
	}

	// Create server (IMPORTANT for graceful shutdown)
	server := &http.Server{
		Addr:    ":" + fmt.Sprint(cfg.Port),
		Handler: mux,
	}

	// Start server in goroutine
	go func() {
		log.Println("Gateway running on", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Listen for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // Ctrl+C (SIGINT)
	// NOTE: os.Kill cannot be caught → don't use it

	sig := <-quit
	log.Printf("Signal %v received. Shutting down...", sig)

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Gateway exited gracefully")
}
