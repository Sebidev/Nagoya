package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: 1, Username: "john_doe", Email: "john@example.com"},
	{ID: 2, Username: "jane_smith", Email: "jane@example.com"},
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the users slice to JSON
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON request body into a User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Validate and process the user data (e.g., save to database)
	// ...

	// Encode the created user as JSON and send it in the response
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func demoapi() {
	// Define the HTTP route and handler
	http.HandleFunc("/get-users", getUsersHandler)
	http.HandleFunc("/users", createUserHandler)

	// Start the HTTP server on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
