package services

import (
	"github.com/rs/zerolog/log"
	"pgtest.com/m/v2/internal/models"
	"pgtest.com/m/v2/internal/repositories"
)

type PlaceService struct {
	placeRepo repositories.PlaceRepository
}

func NewPlaceService(placeRepo repositories.PlaceRepository) *PlaceService {
	return &PlaceService{placeRepo: placeRepo}
}

func (s *PlaceService) CreatePlace(name string, longitude, latitude float64, userID int) (int, error) {
	place := &models.Place{
		Name:      name,
		Longitude: longitude,
		Latitude:  latitude,
		UserID:    userID,
	}

	placeID, err := s.placeRepo.CreatePlace(place)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create place")
		return 0, err
	}
	log.Info().Int("placeID", placeID).Msg("Place created")
	return placeID, nil
}
