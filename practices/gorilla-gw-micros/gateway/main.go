package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Auth service routes, Match all methods (GET, POST, etc.)
	r.HandleFunc("/api/auth/{path:.*}", reverseProxy("http://auth-service:8081", "/api/"))
	// r.HandleFunc("/api/auth/{path:.*}", reverseProxy("http://auth-service:8081", "/api/")).Methods("POST")

	// Products service routes
	r.HandleFunc("/api/products/{path:.*}", reverseProxy("http://product-service:8082", "/api/products/"))

	log.Println("API Gateway running on port 8083...")
	http.ListenAndServe(":8083", r)
}

func reverseProxy(target string, prefix string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := vars["path"]       // Captured path variable (empty for /api/products)
		r.URL.Path = prefix + path // Rewrite the path with the prefix
		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	}
}
