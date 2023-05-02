package main

import (
	"log"

	"github.com/ramseyjiang/go_mid_to_senior/practices/distributed/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")

	log.Fatal(srv.ListenAndServe())
}
