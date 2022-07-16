package main

import (
	"errors"
	"log"
	"net"
	"strconv"

	ss "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/serverstream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	ss.PhoneServer
	contacts []*ss.Contact
}

type pageContact struct {
	CurrentPage int32
	contacts    []*ss.Contact
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50056")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")
	s := grpc.NewServer()
	ss.RegisterPhoneServer(s, &phoneServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (p *phoneServer) AllContacts(_ *ss.AllContactsRequest, stream ss.Phone_AllContactsServer) (err error) {
	p.loadContacts()
	for _, contact := range p.contacts {
		res := &ss.AllContactsResponse{
			Contact: contact,
		}

		// Use the stream object to send the response stream message
		if err = stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func (p *phoneServer) PageContacts(req *ss.PageContactsRequest, stream ss.Phone_PageContactsServer) (err error) {
	p.loadContacts()
	pageSize, _ := strconv.ParseInt(req.PageSize, 10, 64)
	reqPageSize := int32(pageSize)
	if reqPageSize < 1 {
		return errors.New("failed to serve, input page size should be a number and gather than 0")
	}

	pageAmount := (int32(len(p.contacts)) + reqPageSize - 1) / reqPageSize
	pageContent := make([]pageContact, pageAmount)
	for i := 0; int32(i) < pageAmount; i++ {
		pageContent[i].CurrentPage = int32(i + 1)
		pageLastElementIndex := reqPageSize * int32(i+1)
		if reqPageSize*int32(i+1) > int32(len(p.contacts)) {
			pageLastElementIndex = int32(len(p.contacts))
		}
		pageContent[i].contacts = p.contacts[reqPageSize*int32(i) : pageLastElementIndex]
		res := &ss.PageContactsResponse{
			PageAmount:  pageAmount,
			PageSize:    reqPageSize,
			CurrentPage: int32(i + 1),
			Contact:     pageContent[i].contacts,
		}

		// Use the stream object to send the response stream message
		if err = stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

// loadContacts can be defined by yourself or read from database.
func (p *phoneServer) loadContacts() {
	p.contacts = []*ss.Contact{
		0: {
			Firstname: "Stephen",
			Lastname:  "Curry",
			Number:    220123621,
		},
		1: {
			Firstname: "Klay",
			Lastname:  "Thompson",
			Number:    220123632,
		},
		2: {
			Firstname: "Draymond",
			Lastname:  "Green",
			Number:    220123651,
		},
		3: {
			Firstname: "Andrew",
			Lastname:  "Wiggins",
			Number:    220123662,
		},
		4: {
			Firstname: "Jorden",
			Lastname:  "Poole",
			Number:    220123671,
		},
		5: {
			Firstname: "Kevon",
			Lastname:  "Looney",
			Number:    220123621,
		},
		6: {
			Firstname: "Otto",
			Lastname:  "Porter",
			Number:    220123232,
		},
		7: {
			Firstname: "Garry",
			Lastname:  "Payton",
			Number:    220123355,
		},
	}
}
