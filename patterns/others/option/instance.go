package option

import (
	"log"
	"time"
)

type Server struct {
	port    string
	timeout time.Duration
	maxConn int64
}

// Option is the keypoint for the option pattern. It is used to define a type of function which is optional.
type Option func(*Server)

func NewServer(port string, options ...Option) *Server {
	server := &Server{
		port: port,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithMaxConn(maxConn int64) Option {
	return func(s *Server) {
		s.maxConn = maxConn
	}
}

func Trigger() {
	server1 := NewServer("4000")
	log.Println(server1)

	server2 := NewServer("5000", WithTimeout(2*time.Second))
	log.Println(server2)

	server3 := NewServer("6000", WithMaxConn(100))
	log.Println(server3)
}
