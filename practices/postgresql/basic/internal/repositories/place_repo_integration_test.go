package repositories

import (
	"testing"

	_ "github.com/lib/pq" // 确保导入 postgres 驱动
)

// func setupDB(t *testing.T) *sql.DB {
// 	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/mydb?sslmode=disable")
// 	assert.NoError(t, err)
// 	return db
// }

func TestPlaceRepository_CreatePlace(t *testing.T) {
	// db := setupDB(t)
	// repo := NewPlaceRepository(db)
	//
	// place := &models.Place{Name: "Eiffel Tower", Longitude: 2.2945, Latitude: 48.8584, UserID: 1}
	// placeID, err := repo.CreatePlace(place)
	// assert.NoError(t, err)
	// assert.NotZero(t, placeID)
}

func TestPlaceRepository_GetPlaceByID(t *testing.T) {
	// db := setupDB(t)
	// repo := NewPlaceRepository(db)
	//
	// place := &models.Place{Name: "Eiffel Tower", Longitude: 2.2945, Latitude: 48.8584, UserID: 1}
	// placeID, err := repo.CreatePlace(place)
	// assert.NoError(t, err)
	//
	// retrievedPlace, err := repo.GetPlaceByID(placeID)
	// assert.NoError(t, err)
	// assert.Equal(t, place.Name, retrievedPlace.Name)
	// assert.Equal(t, place.Longitude, retrievedPlace.Longitude)
	// assert.Equal(t, place.Latitude, retrievedPlace.Latitude)
}
