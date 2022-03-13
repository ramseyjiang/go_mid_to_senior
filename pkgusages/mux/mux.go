package mucpkg

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Trigger() {
	// The router is the main router for your web application and will later be passed as parameter to the server. It will receive all HTTP connections and pass it on to the request handlers you will register on it.
	r := mux.NewRouter()

	// http://localhost:9009/books/kobe/page/1
	// http://localhost:9009/books/go-programming-blueprint/page/10
	// register request handlers like usual. The only difference is that instead of calling http.HandleFunc(...), you call HandleFunc on your router like this: r.HandleFunc(...).
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// The package comes with the function mux.Vars(r) which takes the http.Request as parameter and returns a map of the segments.
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"]   // the page

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	_ = http.ListenAndServe(":9009", r)

	/*
		Features of the gorilla/mux Router
		Methods
		Restrict the request handler to specific HTTP methods.

		r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
		r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
		r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
		r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")


		Hostnames & Subdomains
		Restrict the request handler to specific hostnames or subdomains.

		r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")


		Schemes
		Restrict the request handler to http/https.

		r.HandleFunc("/secure", SecureHandler).Schemes("https")
		r.HandleFunc("/insecure", InsecureHandler).Schemes("http")


		Path Prefixes & Subrouters
		Restrict the request handler to specific path prefixes.

		bookrouter := r.PathPrefix("/books").Subrouter()
		bookrouter.HandleFunc("/", AllBooks)
		bookrouter.HandleFunc("/{title}", GetBook)
	*/
}
