package pqgorm

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {
	db, err := connectToDatabase()
	require.NoError(t, err, "Failed to connect to the database")

	// Begin a transaction
	tx := db.Begin()
	require.NoError(t, tx.Error, "Failed to begin transaction")

	// Insert a user within the transaction
	user := User{Username: gofakeit.Username(), Email: gofakeit.Email()}

	result := tx.Create(&user)
	require.NoError(t, result.Error, "Failed to insert user")
	fmt.Println(user.ID)

	// Rollback the transaction
	err = tx.Rollback().Error
	require.NoError(t, err, "Failed to rollback transaction")
}

func TestQueryUser(t *testing.T) {
	db, err := connectToDatabase()
	require.NoError(t, err, "Failed to connect to the database")

	// Begin a transaction for inserting a user
	tx1 := db.Begin()
	require.NoError(t, tx1.Error, "Failed to begin transaction")

	// Insert a user within the transaction
	user := User{Username: gofakeit.Username(), Email: gofakeit.Email()}

	result := tx1.Create(&user)
	require.NoError(t, result.Error, "Failed to insert user")
	fmt.Println(user.ID)

	// Commit the transaction
	err = tx1.Commit().Error
	require.NoError(t, err, "Failed to commit transaction")

	// Begin a transaction for querying the user
	tx2 := db.Begin()
	require.NoError(t, tx2.Error, "Failed to begin transaction")

	// Query the user within the transaction
	var queriedUser User
	err = tx2.First(&queriedUser, "id = ?", user.ID).Error
	require.NoError(t, err, "Failed to query user")
	require.Equal(t, user.ID, queriedUser.ID, "Queried user ID does not match the inserted user ID")
	fmt.Println(user.ID)

	// Rollback the transaction
	err = tx2.Rollback().Error
	require.NoError(t, err, "Failed to rollback transaction")
}
