package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/products/list", GetProducts).Methods("GET")
	log.Println("Product service running on port 8082...")
	http.ListenAndServe(":8082", r)
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "iPhone 14", Price: "999 USD"},
		{ID: 2, Name: "MacBook Pro", Price: "1999 USD"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
