package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	run "github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	grpcAddr := ":50100"
	httpAddr := ":8080"

	// Setup gRPC servers.
	baseGrpcServer := grpc.NewServer()
	userGrpcServer := NewUserGRPCServer()
	apiv1.RegisterUserServiceServer(baseGrpcServer, userGrpcServer)

	// Setup gRPC gateway.
	ctx := context.Background()
	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	{
		err := apiv1.RegisterUserServiceHandlerServer(ctx, rmux, userGrpcServer)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Serve.
	var g run.Group
	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			log.Printf("Serving grpc address %s", grpcAddr)
			return baseGrpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			log.Printf("Serving http address %s", httpAddr)
			return http.Serve(httpListener, mux)
		}, func(err error) {
			httpListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}

type userServer struct {
	m map[string]*apiv1.User
}

func NewUserGRPCServer() apiv1.UserServiceServer {
	return &userServer{
		m: map[string]*apiv1.User{},
	}
}

func (s *userServer) CreateUser(ctx context.Context, req *apiv1.CreateUserRequest) (*apiv1.CreateUserResponse, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil,
			status.Error(codes.Internal, err.Error())
	}
	s.m[id.String()] = req.User
	return &apiv1.CreateUserResponse{
		Id: id.String(),
	}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *apiv1.GetUserRequest) (*apiv1.GetUserResponse, error) {
	foundUser, ok := s.m[req.Id]
	if !ok {
		return nil,
			status.Error(codes.NotFound, fmt.Errorf("User not found by id %v", req.Id).Error())
	}
	return &apiv1.GetUserResponse{
		User: &apiv1.UserList{
			Id:   req.Id,
			Name: foundUser.Firstname,
		},
	}, nil
}
