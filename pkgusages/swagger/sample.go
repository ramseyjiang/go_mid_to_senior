package main

import (
	"encoding/json"
	"net/http"
)

// GreetingResponse defines the response structure for the /greeting endpoint
// swagger:response greetingResponse
type GreetingResponse struct {
	// The greeting message
	// in: body
	Message string `json:"message"`
}

// swagger:route GET /greeting greeting getGreeting
//
// # Get a greeting message
//
// Returns a simple greeting message.
//
// Produces:
// - application/json
//
// Responses:
//
//	200: greetingResponse
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	response := &GreetingResponse{
		Message: "Hello, welcome to My Simple API!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Serve the static Swagger UI files
	fs := http.FileServer(http.Dir("swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	// Serve the Swagger YAML file
	http.Handle("/swagger.yml", http.FileServer(http.Dir(".")))

	// Your existing handlers
	http.HandleFunc("/greeting", greetingHandler)
	http.ListenAndServe(":8080", nil)
}
