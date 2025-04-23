package services

import (
	"context"

	userGRPC "pgtest.com/m/v2/internal/services/proto"
)

type UserGRPCService struct {
	userService *UserService
}

func NewUserGRPCService(userService *UserService) *UserGRPCService {
	return &UserGRPCService{userService: userService}
}

func (s *UserGRPCService) CreateUser(ctx context.Context, req *userGRPC.CreateUserRequest) (*userGRPC.CreateUserResponse, error) {
	userID, err := s.userService.CreateUser(req.Name, req.Age)
	if err != nil {
		return nil, err
	}

	return &userGRPC.CreateUserResponse{UserId: int32(userID)}, nil
}

func (s *UserGRPCService) GetUserByID(ctx context.Context, req *userGRPC.GetUserByIDRequest) (*userGRPC.GetUserByIDResponse, error) {
	user, err := s.userService.GetUserByID(int(req.UserId))
	if err != nil {
		return nil, err
	}

	return &userGRPC.GetUserByIDResponse{
		User: &userGRPC.User{
			Id:   int32(user.ID),
			Name: user.Name,
			Age:  int32(user.Age),
		},
	}, nil
}
