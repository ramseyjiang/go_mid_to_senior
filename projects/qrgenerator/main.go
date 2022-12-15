package main

import (
	"image/png"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"fmt"
	"os"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QrText struct {
	Text string `json:"text"`
}

type Page struct {
	Title string
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/generator/", viewCodeHandler).Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	port := "12345"
	log.Println(http.ListenAndServe(":"+port, loggedRouter))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, err := template.ParseFiles("generator.html")
	if err != nil {
		log.Println("Problem parsing html file")
	}

	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")

	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		fmt.Println(err)
	} else {
		qrCode, err = barcode.Scale(qrCode, 128, 128)
		if err != nil {
			fmt.Println(err)
		} else {
			png.Encode(w, qrCode)
		}
	}
}
