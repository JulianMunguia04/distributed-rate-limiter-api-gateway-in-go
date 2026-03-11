package main

import (
	"log"
	"net/http"

	"gateway/internal/loadbalancer"
	"gateway/internal/middleware"
	"gateway/internal/proxy"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env:", err)
	} else {
		log.Println(".env loaded successfully")
	}

	app1LB := loadbalancer.NewRoundRobin([]string{
		"http://localhost:5000",
		"http://localhost:5001",
		"http://localhost:5002",
	})

	app2LB := loadbalancer.NewRoundRobin([]string{
		"http://localhost:3000",
		"http://localhost:3001",
	})

	// reverse proxy
	app1Handler := proxy.ProxyRequest(app1LB)
	app2Handler := proxy.ProxyRequest(app2LB)

	// middleware for app1 (no auth)
	app1Middleware := middleware.Chain(
		app1Handler,
		middleware.Logging,
		middleware.Recovery,
		middleware.RateLimit,
	)

	app2Middleware := middleware.Chain(
		app2Handler,
		middleware.Logging,
		middleware.Recovery,
		middleware.APIKeyAuth,
		middleware.RateLimit,
	)

	http.Handle("/app1/", app1Middleware)
	http.Handle("/app2/", app2Middleware)

	log.Println("Gateway running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
