package nethttp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// It is the net/http package usage.
// Access http://localhost:9009/

func Trigger() {
	baseUsage()
}

// http://localhost:9000/
func baseUsage() {
	// The "HandleFunc" method accepts a path and a function as arguments
	http.HandleFunc("/", handler)

	// After defining our server, we finally "listen and serve" on port 9000
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	err := http.ListenAndServe(":9010", nil) // Listen for browser requests, and respond to them.
	log.Fatal(err)
}

/*
	First, create a Handler which receives all coming HTTP connections from browsers, HTTP clients or API requests.
	A handler in Go is a function with this signature: func (writer http.ResponseWriter, request *http.Request)

	http.ResponseWriter which is where you write your text/html response to.
	http.Request which contains all information about this HTTP request including things like the URL or header fields.

	Second, Registering a request handler to the default HTTP Server is as simple as this: "http.HandleFunc"
*/
// http://localhost:9010/test
// % curl -v -d "test curl" http://localhost:9010
// In the server, it will output following.
// % go run nethttp.go
// 2022/05/21 22:42:35 hello
// 2022/05/21 22:42:35 Data test curl
func handler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello world\n")
	requestBody, _ := ioutil.ReadAll(request.Body)
	log.Println("hello")
	log.Printf("Data %s\n", requestBody)
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	// request.URL.Path will output the url name after '/', for example: http://localhost:9000/test, it will output "/test"
	// n is the number of bytes written.
	n, _ := fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)  // r.URL.Path[0:] is the same with r.URL.Path.
	m, _ := fmt.Fprintf(writer, "Hello, your route is: %s\n", request.URL.Path[1:]) // it will output "test"
	q, _ := fmt.Fprint(writer, "Welcome to my website")
	fmt.Println(n, m, q)
	// http.FileServer and point it to an url path. For the file server to work properly it needs to know, where to serve files from, using http.FileServer to serve static assets like JavaScript, CSS and images
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
}
