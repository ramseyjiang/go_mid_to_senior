package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Message string
}

var conf = &Config{Message: "Before hot reload"}

func multiSignalHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		log.Println("Signal:", signal.String())
		log.Println("After hot reload")
		conf.Message = "Hot reload has been finished."
	case syscall.SIGINT:
		log.Println("Signal:", signal.String())
		log.Println("Interrupt by Ctrl+C")
		os.Exit(0)
	case syscall.SIGTERM:
		log.Println("Signal:", signal.String())
		log.Println("Process is killed.")
		os.Exit(0)
	default:
		log.Println("Unhandled/unknown signal")
	}
}

func router() {
	log.Println("starting up....")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(conf.Message))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func main() {
	router()
	// When a server shut down, it should stop receiving new requests while completing the ongoing requests,
	// return their responses, and then shut down.
	// create a channel for listening to OS signals, it needs to reserve to buffer size 1, so the notifier are not blocked
	sigCh := make(chan os.Signal, 1)

	// signal.Notify will send a signal to sigCh channel when program is interrupted.
	// syscall.SIGINT is used shutdown gracefully on Ctrl+C, it equals os.Interrupt also.
	// syscall.SIGTERM is the usual signal for termination and the default one for docker containers, which is also used by Kubernetes.
	// syscall.SIGHUP is used for the hot reload configuration.
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		multiSignalHandler(<-sigCh)
	}
}
