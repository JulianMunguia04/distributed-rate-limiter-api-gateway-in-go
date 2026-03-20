package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var counter int
var mu sync.Mutex

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run server.go <port> <delay_in_seconds>")
	}

	port := os.Args[1]

	// Convert delay argument to integer
	delaySeconds, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Delay must be a number (seconds)")
	}

	delay := time.Duration(delaySeconds) * time.Second

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Ignore favicon
		if r.URL.Path == "/favicon.ico" {
			http.NotFound(w, r)
			return
		}

		// Delay
		time.Sleep(delay)

		mu.Lock()
		counter++
		//current := counter
		mu.Unlock()

		//fmt.Fprintf(w, "Instance running on port %s | Count: %d | Delay: %ds\n", port, current, delaySeconds)
		fmt.Fprintf(w, "Running on port %s", port)
	})

	//Test health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("Starting server on port", port, "with delay", delaySeconds, "seconds")

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

//go run server.go <port number> <delay>

// or go build -o server.exe server.go
// then
// .\server.exe <port number> <delay>
