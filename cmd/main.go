package main

import (
	"log"
	"net/http"
	"os"
)

func main{

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handler endpoints
	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.LOCATION_PATH, handler.LocationHandler)
	http.HandleFunc(handler.COLLECTION_PATH, handler.CollectionHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}