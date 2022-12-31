package main

import (
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
		log.Println("Signal:", signal.String())
		log.Println("hot reload")
		conf.Message = "Go To Hell!"
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

func routerExample() {
	log.Println("starting up....")
	// Output is used to display hot reload result.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(conf.Message))
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

func main() {
	routerExample()
	// When a server shut down, it should stop receiving new requests while complete the ongoing requests,
	// and return their responses, and then, shut down finally.
	// create a channel for listening to OS signals, it needs to reserve to buffer size 1, so the notifier are not blocked
	sigCh := make(chan os.Signal, 1)

	// signal.Notify will send a signal to sigCh channel when program is interrupted.
	// SIGINT is used shutdown gracefully on Ctrl+C, it equals to os.Interrupt also.
	// syscall.SIGTERM is the usual signal for termination and the default one for docker containers, which is also used by kubernetes.
	// syscall.SIGHUP is used for the signal to reload configuration.
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		s := <-sigCh
		multiSignalHandler(s)
	}
}
