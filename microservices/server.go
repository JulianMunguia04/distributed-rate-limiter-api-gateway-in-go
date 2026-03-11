package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a port number. Example: go run server.go 5000")
	}

	port := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Ignore favicon
		if r.URL.Path == "/favicon.ico" {
			http.NotFound(w, r)
			return
		}

		mu.Lock()
		counter++
		current := counter
		mu.Unlock()

		fmt.Fprintf(w, "Instance running on port %s | Count: %d\n", port, current)
	})

	log.Println("Starting server on port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

//go run server.go <port number>
