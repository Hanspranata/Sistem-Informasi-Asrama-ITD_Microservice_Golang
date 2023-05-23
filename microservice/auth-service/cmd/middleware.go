package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new middleware handler
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Perform middleware logic here
			log.Println("Executing middleware")

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}

	// Create a new HTTP server with the middleware
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware(http.DefaultServeMux),
	}

	// Start the server
	log.Println("Server started on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
