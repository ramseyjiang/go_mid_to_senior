package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func multiSignalHandler(signal os.Signal) {
	switch signal {
	case syscall.SIGHUP:
		fmt.Println("Signal:", signal.String())
		os.Exit(0)
	case syscall.SIGINT:
		fmt.Println("Signal:", signal.String())
		os.Exit(0)
	case syscall.SIGTERM:
		fmt.Println("Signal:", signal.String())
		os.Exit(0)
	case syscall.SIGQUIT:
		fmt.Println("Signal:", signal.String())
		os.Exit(0)
	default:
		fmt.Println("Unhandled/unknown signal")
	}
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	exitCh := make(chan int)

	fmt.Println("start")
	go func() {
		for {
			s := <-sigCh
			multiSignalHandler(s)
		}
	}()

	exitCode := <-exitCh
	os.Exit(exitCode)
}
