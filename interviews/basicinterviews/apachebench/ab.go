package main

import (
	"fmt"

	"net/http"

	"time"
)

const maxConcurrentRequests = 200

// Create a request pool to limit the number of concurrent requests

var sem = make(chan struct{}, maxConcurrentRequests)

// Function to handle each HTTP request

func handleRequest(w http.ResponseWriter, r *http.Request) {

	// Get the goroutine signal

	sem <- struct{}{}

	defer func() {

		<-sem

	}()

	// Simulate request processing logic

	time.Sleep(time.Millisecond * 10) // Simulate processing time

	fmt.Fprintf(w, "Hello, World!")

}

// After start the server,
// Run anywhere using: ab -n 5000 -c 140 http://localhost:8080/, do the benchmark test
func main() {
	http.HandleFunc("/", handleRequest)

	// Start the HTTP server
	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
