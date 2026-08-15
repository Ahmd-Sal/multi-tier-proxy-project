package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	// Log the request
	log.Printf("[%s] %s requested from %s", r.Method, r.URL.Path, r.RemoteAddr)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error retrieving hostname: %v", err)
		http.Error(w, "Unable to get hostname", http.StatusInternalServerError)
		return
	}

	// Send the hostname back in the HTTP response
	fmt.Fprintf(w, "Container Hostname: %s\n", hostname)
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s requested from %s", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "This is the Health endpoint /health")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/health", getHealth)

	// Log when the server starts
	log.Println("Server is running on port :80...")

	// Starts listening on port 80
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}

}
