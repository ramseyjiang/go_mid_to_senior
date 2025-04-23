package services

import (
	"pgtest.com/m/v2/internal/models"
	"pgtest.com/m/v2/internal/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(name string, age int32) (int, error) {
	user := &models.User{Name: name, Age: int(age)}
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}
