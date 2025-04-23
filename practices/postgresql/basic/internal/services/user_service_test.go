package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pgtest.com/m/v2/internal/models"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateTable() error {
	// TODO implement me
	panic("implement me")
}

func (m *MockUserRepository) CreateUser(user *models.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id int) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{Name: "张三", Age: 30}
	mockRepo.On("CreateUser", user).Return(1, nil)

	userID, err := service.CreateUser("张三", 30)
	assert.NoError(t, err)
	assert.Equal(t, 1, userID)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	user := &models.User{ID: 1, Name: "张三", Age: 30}
	mockRepo.On("GetUserByID", 1).Return(user, nil)

	retrievedUser, err := service.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)

	mockRepo.AssertExpectations(t)
}
