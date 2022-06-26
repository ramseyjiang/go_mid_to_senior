package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"

	cal "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unary/tlscal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type TlsCalService struct {
	cal.CalculateServiceServer
}

func (s *TlsCalService) Add(ctx context.Context, request *cal.Request) (*cal.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b
	return &cal.Response{Result: result}, nil
}

func (s *TlsCalService) Sub(ctx context.Context, request *cal.Request) (*cal.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b
	return &cal.Response{Result: result}, nil
}

func (s *TlsCalService) Mul(ctx context.Context, request *cal.Request) (*cal.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b
	return &cal.Response{Result: result}, nil
}

func main() {
	// listen port
	listen, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("list port err: %v", err)
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))

	// register service into grpc server
	cal.RegisterCalculateServiceServer(grpcServer, &TlsCalService{})
	// reflection.Register(grpcServer) // TODO: try to make grpcui works in tls

	log.Printf("listening at %v", listen.Addr())

	// listen port
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("grpc serve err: %v", err)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// read ca's cert, verify to client's certificate
	pemCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// read server cert & key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// config the certificate
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
