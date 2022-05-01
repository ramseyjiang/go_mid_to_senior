package main

import (
	"log"
	"net"
)

func main() {
	// Open a TCP Session, use the net.Dial() method to open a TCP connection to the same localhost:9090 address our TCP server is listening on.
	c, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Unable to open TCP Connection: %s", err)
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(c)

	// With the returned net.Conn is writing our sample message using the net.Conn.Write() method.
	log.Printf("TCP Session Open")
	b := []byte("Hello, is there anybody out there?")
	_, err = c.Write(b)
	if err != nil {
		log.Printf("Error from TCP Session: %s", err)
	}

	// Read any responses until we get an error
	// looping a net.Conn.Read() method to continuously read data sent from the TCP server.
	for {
		d := make([]byte, 120)
		_, err := c.Read(d)
		if err != nil {
			log.Fatalf("Error from TCP Session: %s", err)
		}
	}
}
