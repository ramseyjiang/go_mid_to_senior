package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// Create a listener, Go is telling the system kernel to bind port 9000 on all of available interfaces.
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Printf("Listener returned: %s", err)
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(l)

	for {
		// Accept new connections, l.Accept() is equal to net.Listener.Accept() method.
		// The Accept() method will wait until a new connection arrives and returns that connection as a net.Conn.
		c, err := l.Accept()
		if err != nil {
			log.Printf("Unable to accept new connections: %s", err)
		}

		// Create a goroutine that reads and writes-back data
		// this goroutine takes the net.Conn referred to as c and starts reading and writing to the connection.
		go func() {
			log.Printf("TCP Session Open")
			// Clean up session when goroutine completes, it's ok to
			// call Close more than once.
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

		// Create another goroutine that closes a session after 15 seconds.
		// this goroutine is using time.After to wait for 15 seconds.
		// Once those 15 seconds are over, our goroutine will call net.Conn.Close() via the defer function.
		go func() {
			<-time.After(time.Duration(15) * time.Second)
			defer func(c net.Conn) {
				err := c.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(c)
		}()
	}
}
