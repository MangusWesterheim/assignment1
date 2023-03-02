package main

import (
	"log"
	"net/http"
	"os"
)

func main{
// Extract PORT variable from the environment variables
port := os.Getenv("PORT")

// Override port with default port if not provided (e.g. local deployment)
if port == "" {
log.Println("$PORT has not been set. Default: 8080")
port = "8080"
}

// Default handler for requests (just displays information and points to /diag)
http.HandleFunc("/", defaultHandler)
// Assign path for diagnostics handler (actual service feature)
http.HandleFunc("/diag", diagHandler)

// Start HTTP server
log.Println("Starting server on port " + port + " ...")
log.Fatal(http.ListenAndServe(":"+port, nil))
}