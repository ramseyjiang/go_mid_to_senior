package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	tlsunary "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/alltypes/tlsunary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type phoneServer struct {
	tlsunary.PhoneServer
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		_ = errors.New("failed to listen: the port")
	}
	log.Print("Server started")

	optServer := genServerTLSOpts()
	s := grpc.NewServer(optServer...)
	tlsunary.RegisterPhoneServer(s, &phoneServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func genServerTLSOpts() []grpc.ServerOption {
	var opts []grpc.ServerOption
	certFile := "cert/server.crt"
	keyFile := "cert/server.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed loading certficates: %v\n", err)
	}
	opts = append(opts, grpc.Creds(creds))
	return opts
}

func (p *phoneServer) GetContactName(ctx context.Context, req *tlsunary.GetContactNameRequest) (resp *tlsunary.GetContactNameResponse, err error) {
	reqNum, _ := strconv.ParseUint(req.Number, 10, 64)

	if len(req.Number) < 9 {
		return &tlsunary.GetContactNameResponse{}, errors.New("invalid number")
	}

	contacts := p.loadContacts()
	for i := 0; i < len(contacts); i++ {
		if reqNum == contacts[i].Number {
			resp = &tlsunary.GetContactNameResponse{
				Firstname: contacts[i].Firstname,
				Lastname:  contacts[i].Lastname,
			}
			return resp, nil
		}
	}

	return resp, errors.New("no contact found")
}

func (p *phoneServer) GetContactNum(ctx context.Context, req *tlsunary.GetContactNumRequest) (resp *tlsunary.GetContactNumResponse, err error) {
	if req.Firstname == "" || req.Lastname == "" {
		return &tlsunary.GetContactNumResponse{}, errors.New("invalid firstname and lastname input")
	}

	contacts := p.loadContacts()
	for i := 0; i < len(contacts); i++ {
		if req.Firstname == contacts[i].Firstname && req.Lastname == contacts[i].Lastname {
			resp = &tlsunary.GetContactNumResponse{
				Num:    contacts[i].Number,
				Result: "Owner is " + contacts[i].Firstname + " " + contacts[i].Lastname,
			}
			return resp, nil
		}
	}

	return resp, errors.New("no contact found")
}

func (p *phoneServer) ListContacts(context.Context, *tlsunary.ListContactsRequest) (resp *tlsunary.ListContactsResponse, err error) {
	contacts := p.loadContacts()

	return &tlsunary.ListContactsResponse{
		Contacts: contacts,
		Sum:      int32(len(contacts)),
	}, nil
}

// loadContacts can be defined by yourself or read from database.
func (p *phoneServer) loadContacts() (contacts []*tlsunary.Contacts) {
	contacts = []*tlsunary.Contacts{
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
