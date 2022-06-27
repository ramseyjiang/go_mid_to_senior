package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	unary "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	unary.PhoneServer
}

const numMinLen = 9

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")
	s := grpc.NewServer()
	unary.RegisterPhoneServer(s, &phoneServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (p *phoneServer) GetContactName(ctx context.Context, req *unary.GetContactNameRequest) (resp *unary.GetContactNameResponse, err error) {
	reqNum, _ := strconv.ParseUint(req.Number, 10, 64)

	if len(req.Number) < numMinLen {
		return &unary.GetContactNameResponse{}, errors.New("invalid number")
	}

	contacts := p.loadContacts()
	for i := 0; i < len(contacts); i++ {
		if reqNum == contacts[i].Number {
			resp = &unary.GetContactNameResponse{
				Firstname: contacts[i].Firstname,
				Lastname:  contacts[i].Lastname,
			}
			return resp, nil
		}
	}

	return resp, errors.New("no contact found")
}

func (p *phoneServer) GetContactNum(ctx context.Context, req *unary.GetContactNumRequest) (resp *unary.GetContactNumResponse, err error) {
	if req.Firstname == "" || req.Lastname == "" {
		return &unary.GetContactNumResponse{}, errors.New("invalid firstname and lastname input")
	}

	contacts := p.loadContacts()
	for i := 0; i < len(contacts); i++ {
		if req.Firstname == contacts[i].Firstname && req.Lastname == contacts[i].Lastname {
			resp = &unary.GetContactNumResponse{
				Num:    contacts[i].Number,
				Result: "Owner is " + contacts[i].Firstname + " " + contacts[i].Lastname,
			}
			return resp, nil
		}
	}

	return resp, errors.New("no contact found")
}

func (p *phoneServer) ListContacts(context.Context, *unary.ListContactsRequest) (resp *unary.ListContactsResponse, err error) {
	contacts := p.loadContacts()

	return &unary.ListContactsResponse{
		Contacts: contacts,
		Sum:      int32(len(contacts)),
	}, nil
}

// loadContacts can be defined by yourself or read from database.
func (p *phoneServer) loadContacts() (contacts []*unary.Contacts) {
	contacts = []*unary.Contacts{
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
	return
}
