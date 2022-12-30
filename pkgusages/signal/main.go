package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var conf = &Config{Message: "Hello, World!"}

type Config struct {
	Message string
}

func multiSignalHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		fmt.Println("Signal:", signal.String())
		log.Println("hot reload")
		conf.Message = "Go To Hell!"
	case syscall.SIGINT:
		fmt.Println("Signal:", signal.String())
	case syscall.SIGTERM:
		fmt.Println("Signal:", signal.String())
	case syscall.SIGQUIT:
		fmt.Println("Signal:", signal.String())
	default:
		fmt.Println("Unhandled/unknown signal")
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(conf.Message))
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("starting up....")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	for {
		s := <-sigCh
		multiSignalHandler(s)
	}

	// for {
	// 	select {
	// 	case <-sigCh:
	// 		log.Println("hot reload")
	// 		conf.Message = "Go To Hell!"
	// 	}
	// }
}
