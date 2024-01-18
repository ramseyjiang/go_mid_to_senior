package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/connect", contactHandler)

	http.HandleFunc("/user", userHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("POST request to the Contact Page!"))
	case "GET":
		w.Write([]byte("GET request to the Contact Page!"))
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	w.Write([]byte("User ID: " + userID))
}
