package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a file server handler that serves files from the specified folder
	fileServer := http.FileServer(http.Dir("./nagoya-frontend/dist/"))

	// Register the file server handler to handle all requests
	http.Handle("/", fileServer)

	// Start the HTTP server on port 8080
	log.Println("Server started. Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
