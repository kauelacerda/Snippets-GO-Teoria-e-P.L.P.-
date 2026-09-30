package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint reached: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/health", handler)
	// Starts a robust HTTP server directly from standard library
	http.ListenAndServe(":8080", nil)
}
