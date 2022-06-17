package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	cal "github.com/ramseyjiang/go_mid_to_senior/pkgusages/grpc/unary/tlscal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	tlsCredential, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.Dial(
		"0.0.0.0:50053",
		grpc.WithTransportCredentials(tlsCredential),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	// Contact the server and print out its response.
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := cal.NewCalculateServiceClient(conn)
	doAdd(client)
	doSub(client)
	doMul(client)
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// read ca's cert
	pemCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	// read server cert
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func doAdd(client cal.CalculateServiceClient) {
	addResp, err := client.Add(context.Background(), &cal.Request{A: 1, B: 2})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Add result is ", addResp.Result)
}

func doSub(client cal.CalculateServiceClient) {
	subResp, err := client.Add(context.Background(), &cal.Request{A: 5, B: 3})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Sub result is ", subResp.Result)
}

func doMul(client cal.CalculateServiceClient) {
	mulResp, err := client.Mul(context.Background(), &cal.Request{A: 5, B: 3})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Mul result is ", mulResp.Result)
}
