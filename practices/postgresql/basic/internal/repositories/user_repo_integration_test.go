package repositories

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"pgtest.com/m/v2/internal/models"
)

func setupDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "postgres://user:password@postgres_container:5432/mydb?sslmode=disable")
	assert.NoError(t, err)
	return db
}

func TestUserRepository_CreateUser(t *testing.T) {
	db := setupDB(t)
	repo := NewUserRepository(db)

	user := &models.User{Name: "张三", Age: 30}
	userID, err := repo.CreateUser(user)
	assert.NoError(t, err)
	assert.NotZero(t, userID)
}

func TestUserRepository_GetUserByID(t *testing.T) {
	db := setupDB(t)
	repo := NewUserRepository(db)

	user := &models.User{Name: "张三", Age: 30}
	userID, err := repo.CreateUser(user)
	assert.NoError(t, err)

	retrievedUser, err := repo.GetUserByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, retrievedUser.Name)
	assert.Equal(t, user.Age, retrievedUser.Age)
}
