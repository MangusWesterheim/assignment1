package main

import (
	"log"
	"net/http"
	"os"
	assignment1 "prog2005"
)

func main() {

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handler endpoints
	http.HandleFunc(assignment1.DEFAULT_PATH, assignment1.EmptyHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
