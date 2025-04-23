package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pgtest.com/m/v2/internal/models"
)

type MockPlaceRepository struct {
	mock.Mock
}

func (m *MockPlaceRepository) CreateTable() error {
	// TODO implement me
	panic("implement me")
}

func (m *MockPlaceRepository) GetPlaceByID(id int) (*models.Place, error) {
	// TODO implement me
	panic("implement me")
}

func (m *MockPlaceRepository) UpdatePlace(place *models.Place) error {
	// TODO implement me
	panic("implement me")
}

func (m *MockPlaceRepository) DeletePlace(id int) error {
	// TODO implement me
	panic("implement me")
}

func (m *MockPlaceRepository) CreatePlace(place *models.Place) (int, error) {
	args := m.Called(place)
	return args.Int(0), args.Error(1)
}

func TestPlaceService_CreatePlace(t *testing.T) {
	mockRepo := new(MockPlaceRepository)
	service := NewPlaceService(mockRepo)

	place := &models.Place{Name: "Eiffel Tower", Longitude: 2.2945, Latitude: 48.8584, UserID: 1}
	mockRepo.On("CreatePlace", place).Return(1, nil)

	placeID, err := service.CreatePlace("Eiffel Tower", 2.2945, 48.8584, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, placeID)

	mockRepo.AssertExpectations(t)
}
