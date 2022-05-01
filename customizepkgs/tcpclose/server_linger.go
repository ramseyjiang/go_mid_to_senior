package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// Create a listener
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Listener returned: %s", err)
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(l)

	for {
		// Accept new connections
		c, err := l.Accept()
		if err != nil {
			log.Printf("Unable to accept new connections: %s", err)
		}

		// Create a goroutine that reads and writes-back data
		go func() {
			log.Printf("TCP Session Open")
			// Clean up session when goroutine completes, it's ok to call Close more than once.
			defer func(c net.Conn) {
				err := c.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(c)

			for {
				b := make([]byte, 120)

				// Read from TCP Buffer
				_, err := c.Read(b)
				if err != nil {
					log.Printf("Error reading TCP Session: %s", err)
					break
				}

				// Write-back data to TCP Client
				_, err = c.Write(b)
				if err != nil {
					log.Printf("Error writing TCP Session: %s", err)
					break
				}
			}
		}()

		// Create a goroutine that closes a session after 15 seconds
		go func() {
			<-time.After(time.Duration(15) * time.Second)

			// Use SetLinger to force close the connection, The value in SetLinger passed can be thought of as a timer value in seconds.
			// SetLinger sets the behavior of Close on a connection which still has data waiting to be sent or to be acknowledged.
			err := c.(*net.TCPConn).SetLinger(1)
			if err != nil {
				log.Printf("Error when setting linger: %s", err)
			}

			defer func(c net.Conn) {
				err := c.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(c)
		}()
	}
}
