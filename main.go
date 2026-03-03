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
	fmt.Println("Path:", r.URL.Path)

	mu.Lock()
	counter++
	val := counter
	mu.Unlock()

	fmt.Fprintf(w, "Count: %d\n", val)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {	//Avoid double request
		http.NotFound(w, r)
	})
	http.ListenAndServe(":8080", nil)
}