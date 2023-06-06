package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// User is a struct for a user
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (u User) IsEmpty() bool {
	return u == User{}
}

var users []User
var logServiceURL = "http://localhost:8081/log/add"

// getUsers returns a list of users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// addUser adds a new user
func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if newUser.IsEmpty() {
		http.Error(w, "User cannot be empty", http.StatusBadRequest)
		return
	}

	// Add this check
	if len(newUser.Name) == 0 {
		http.Error(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}

	newUser.ID = uuid.New().String()
	users = append(users, newUser)

	// Create a map with keys "action" and "user"
	logInfo := map[string]interface{}{
		"action": "User added",
		"user":   newUser,
	}

	// Marshal the map to a JSON string
	logMessageJSON, err := json.Marshal(logInfo)
	if err != nil {
		fmt.Printf("Error marshaling log info: %v\n", err)
		return
	}

	logMessage := string(logMessageJSON)
	fmt.Println("Log message:", logMessage)

	// Send the log message to the log microservice
	_, err = http.Post(logServiceURL, "application/json", strings.NewReader(logMessage))
	if err != nil {
		fmt.Printf("Error sending log message: %v\n", err)
	}
}

func main() {
	http.Handle("/user/create", http.HandlerFunc(addUser))
	http.Handle("/user/list", http.HandlerFunc(getUsers))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
