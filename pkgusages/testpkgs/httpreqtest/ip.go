package httpreqtest

import (
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://ipecho.net/plain"

func getIP() string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the body to type string and return
	return string(body)
}
