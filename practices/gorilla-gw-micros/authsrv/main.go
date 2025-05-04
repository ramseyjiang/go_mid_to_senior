package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", LoginHandler).Methods("POST")
	r.HandleFunc("/api/register", RegisterHandler).Methods("GET")
	log.Println("Auth service running on port 8081...")
	http.ListenAndServe(":8081", r)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if user.Username == "admin" && user.Password == "password" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Login successful!"}`))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"error": "Invalid credentials"}`))
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User registered successfully!"}`))
}
