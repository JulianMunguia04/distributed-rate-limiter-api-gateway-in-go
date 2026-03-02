package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	val := counter
	mu.Unlock()

	fmt.Fprintf(w, "Count: %d\n", val)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}