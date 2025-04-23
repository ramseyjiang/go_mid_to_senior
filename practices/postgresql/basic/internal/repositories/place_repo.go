package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"pgtest.com/m/v2/internal/models"
)

type PlaceRepository interface {
	CreateTable() error
	CreatePlace(place *models.Place) (int, error)
	GetPlaceByID(id int) (*models.Place, error)
	UpdatePlace(place *models.Place) error
	DeletePlace(id int) error
}

type PostgresPlaceRepository struct {
	db *sql.DB
}

func NewPlaceRepository(db *sql.DB) PlaceRepository {
	// add PostGIS extension during the initialization
	if _, err := db.Exec("CREATE EXTENSION IF NOT EXISTS postgis; CREATE EXTENSION IF NOT EXISTS postgis_topology;"); err != nil {
		log.Fatal().Err(err).Msg("Failed to enable PostGIS extensions")
	}
	log.Info().Msg("PostGIS extensions enabled successfully!")

	return &PostgresPlaceRepository{db: db}
}

func (r *PostgresPlaceRepository) CreateTable() error {
	// 检查表是否已经存在
	var tableExists bool
	err := r.db.QueryRow("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'places')").Scan(&tableExists)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check if table exists")
		return fmt.Errorf("failed to check if table exists: %w", err)
	}
	if tableExists {
		log.Info().Msg("Table 'places' already exists, skipping creation.")
		return nil
	}

	query := `CREATE TABLE IF NOT EXISTS places (
		id SERIAL PRIMARY KEY,
		name TEXT,
		longitude FLOAT,
		latitude FLOAT,
		user_id INT REFERENCES users(id) ON DELETE CASCADE
	)`

	_, err = r.db.Exec(query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create table")
	}
	return err
}

func (r *PostgresPlaceRepository) CreatePlace(place *models.Place) (int, error) {
	query := `INSERT INTO places (name, longitude, latitude, user_id) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := r.db.QueryRow(query, place.Name, place.Longitude, place.Latitude, place.UserID).Scan(&id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create place")
		return 0, fmt.Errorf("failed to create place: %w", err)
	}
	return id, nil
}

func (r *PostgresPlaceRepository) GetPlaceByID(id int) (*models.Place, error) {
	query := `SELECT id, name, longitude, latitude, user_id FROM places WHERE id = $1`

	place := &models.Place{}
	err := r.db.QueryRow(query, id).Scan(&place.ID, &place.Name, &place.Longitude, &place.Latitude, &place.UserID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get place")
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("place not found")
		}
		return nil, fmt.Errorf("failed to get place: %w", err)
	}
	return place, nil
}

func (r *PostgresPlaceRepository) UpdatePlace(place *models.Place) error {
	query := `UPDATE places SET name = $1, longitude = $2, latitude = $3, user_id = $4 WHERE id = $5`

	_, err := r.db.Exec(query, place.Name, place.Longitude, place.Latitude, place.UserID, place.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update place")
		return fmt.Errorf("failed to update place: %w", err)
	}
	return nil
}

func (r *PostgresPlaceRepository) DeletePlace(id int) error {
	query := `DELETE FROM places WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete place")
		return fmt.Errorf("failed to delete place: %w", err)
	}
	return nil
}
