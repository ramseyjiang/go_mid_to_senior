package gorillaws

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

/**
This client will keep emitting messages after every 1 second.
*/

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func TriggerClient() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done.
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate.

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	socketURL := "ws://localhost:8080" + "/socket"
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}

	defer func(conn *websocket.Conn) {
		err1 := conn.Close()
		if err1 != nil {
			return
		}
	}(conn)
	go receiveHandler(conn)

	// use an infinite loop for listening to events through channels using select.
	// We write a message using conn.WriteMessage() every second.
	// If the interrupt signal is activated, any pending connections are closed, and we exit gracefully!
	// We send our relevant packets here
	for {
		select {
		case <-time.After(time.Duration(1) * time.Millisecond * 1000):
			// Send an echo packet every second
			err2 := conn.WriteMessage(websocket.TextMessage, []byte("Hello from gorilla websocket!"))
			if err2 != nil {
				log.Println("Error during writing to websocket:", err2)
				return
			}
		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully…
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err3 := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err3 != nil {
				log.Println("Error during closing websocket:", err3)
				return
			}

			// If the receiveHandler channel exits, the channel 'done' will be closed. This is the first case <-done condition
			// If the 'done' channel does NOT close, there will be a timeout after 1 second, so the program WILL exit after a second timeout
			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting….")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting….")
			}
			return
		}
	}
}
