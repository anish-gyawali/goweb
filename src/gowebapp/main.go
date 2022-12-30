package main

import (
	"fmt"
	"gowebapp/handlers" // Created own package for handling routes
	"log"
	"net/http"
	"os"
)

// Main program starts here
func main() {
	// Handling Routes
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/task", handlers.TaskHandler)

	// Use environment variable "PORT" or otherwise assign
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("This port is %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)

	// Create server or exit
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
