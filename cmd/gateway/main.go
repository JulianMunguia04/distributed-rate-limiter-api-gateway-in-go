package main

import (
	"log"
	"net/http"

	"gateway/internal/loadbalancer"
	"gateway/internal/proxy"
)

func main() {

	app1LB := loadbalancer.NewRoundRobin([]string{
		"http://localhost:5000",
		"http://localhost:5001",
		"http://localhost:5002",
	})

	app2LB := loadbalancer.NewRoundRobin([]string{
		"http://localhost:3000",
		"http://localhost:3001",
	})

	http.HandleFunc("/app1/", proxy.ProxyRequest(app1LB))
	http.HandleFunc("/app2/", proxy.ProxyRequest(app2LB))

	log.Println("Gateway running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
