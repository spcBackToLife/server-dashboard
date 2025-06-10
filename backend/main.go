package main

import (
	"fmt"
	"log"
	"net/http"
	"wallet-app/backend/handlers/auth" // Import the auth package
)

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Base path handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { // Basic check to ensure only root path matches
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Welcome to the Wallet App Backend!")
	})

	// API v1 routes
	apiV1Mux := http.NewServeMux()
	apiV1Mux.HandleFunc("/auth/register", auth.RegisterHandler)
	apiV1Mux.HandleFunc("/auth/login", auth.LoginHandler)
	// Add other auth handlers here e.g. apiV1Mux.HandleFunc("/auth/login", auth.LoginHandler)

	// Prefix routes for /api/v1
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Mux))


	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
