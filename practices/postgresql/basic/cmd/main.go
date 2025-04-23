package main

import (
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"pgtest.com/m/v2/internal/database"
	"pgtest.com/m/v2/internal/repositories"
	"pgtest.com/m/v2/internal/services"
	userGRPC "pgtest.com/m/v2/internal/services/proto"
)

type userServiceServer struct {
	userGRPC.UnimplementedUserServiceServer
	userGRPCService *services.UserGRPCService
}

func startGRPCServer(userGRPCService *services.UserGRPCService) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	userGRPC.RegisterUserServiceServer(grpcServer, &userServiceServer{userGRPCService: userGRPCService})
	log.Info().Msg("gRPC server is running on port 50051")
	return grpcServer.Serve(lis)
}

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Database connection failed")
	}
	defer db.Close()

	// database migration
	if err := database.RunMigrations(db); err != nil {
		log.Fatal().Err(err).Msg("Failed to apply migrations")
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userGRPCService := services.NewUserGRPCService(userService)

	// Start gRPC server
	if err := startGRPCServer(userGRPCService); err != nil {
		log.Fatal().Err(err).Msg("Failed to start gRPC server")
	}
}
