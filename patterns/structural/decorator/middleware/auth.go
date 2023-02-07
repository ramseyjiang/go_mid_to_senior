package middleware

import (
	"fmt"
	"net/http"
)

// Middleware type is defined as a type alias for a function that takes a http.HandlerFunc and returns a http.HandlerFunc. The LoggingMiddleware and AuthenticationMiddleware functions are examples of middleware that can be used to wrap an HTTP handler to add additional functionality. The HelloHandler function is an HTTP handler that returns "Hello, World!". The ApplyMiddleware function takes an HTTP handler and a slice of middleware functions, and returns a new HTTP handler that applies the middleware functions in the order they appear in the slice. The main function uses ApplyMiddleware to wrap the HelloHandler with the LoggingMiddleware and AuthenticationMiddleware functions, and then registers the resulting HTTP handler with the http.HandleFunc method to serve it on http://localhost:8080.
type Middleware func(http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.Method, r.URL.Path)
		next(w, r)
	}
}

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "secret_token" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func ApplyMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
