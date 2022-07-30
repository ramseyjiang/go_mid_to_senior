package main

import (
	"errors"
	"io"
	"log"
	"net"
	"strconv"

	cs "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/clientstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	cs.PhoneServer
	contacts []*cs.Contact
}

func (p *phoneServer) contains(s []int, key int) bool {
	for _, k := range s {
		if k == key {
			return true
		}
	}
	return false
}

func (p *phoneServer) NumCheck(stream cs.Phone_NumCheckServer) (err error) {
	var results []*cs.Result
	var result *cs.Result
	p.loadContacts()

	for {
		// Start receiving stream messages from the client
		req, err := stream.Recv()

		// Check if the stream has finished
		if errors.Is(err, io.EOF) {
			// Close the connection and return the response to the client
			return stream.SendAndClose(&cs.NumCheckResponse{
				CheckResult: results,
			})
		}

		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
			return err
		}

		// process each request after every err != nil
		k, _ := strconv.Atoi(req.Number)
		if p.contacts[k] != nil {
			result = &cs.Result{
				Msg:    p.contacts[k].Firstname + " " + p.contacts[k].Lastname + " phone is " + strconv.FormatUint(p.contacts[k].PhoneNumber, 10) + ".",
				Status: "Stay",
			}
		} else {
			result = &cs.Result{
				Msg:    req.Number + " Has joined another team.",
				Status: "Left",
			}
		}
		// append a new result
		results = append(results, result)
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50057")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")
	s := grpc.NewServer()
	cs.RegisterPhoneServer(s, &phoneServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// loadContacts can be defined by yourself or read from database.
func (p *phoneServer) loadContacts() {
	p.contacts = []*cs.Contact{
		30: {
			Firstname:   "Stephen",
			Lastname:    "Curry",
			PhoneNumber: 220123621,
		},
		11: {
			Firstname:   "Klay",
			Lastname:    "Thompson",
			PhoneNumber: 220123632,
		},
		23: {
			Firstname:   "Draymond",
			Lastname:    "Green",
			PhoneNumber: 220123651,
		},
		22: {
			Firstname:   "Andrew",
			Lastname:    "Wiggins",
			PhoneNumber: 220123662,
		},
		3: {
			Firstname:   "Jorden",
			Lastname:    "Poole",
			PhoneNumber: 220123671,
		},
		5: {
			Firstname:   "Kevon",
			Lastname:    "Looney",
			PhoneNumber: 220123621,
		},
		33: {
			Firstname:   "James",
			Lastname:    "Wiseman",
			PhoneNumber: 220123623,
		},
	}
}
