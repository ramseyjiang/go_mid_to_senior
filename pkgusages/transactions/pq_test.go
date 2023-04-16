package transactions

import (
	"fmt"
	"log"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestInsertUserCommitted(t *testing.T) {
	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}

	user := User{Username: "Test User", Password: "12345678", Email: gofakeit.Email()}

	err = insertUser(tx, &user)
	assert.NoError(t, err, "Failed to insert user")
	assert.NotZero(t, user.ID, "User1 ID should be non-zero after insertion")
	fmt.Println(user.ID)

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestInsertUserRollback(t *testing.T) {
	db, err := connectToDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Start a transaction
	tx, err := db.Beginx()
	if err != nil {
		log.Fatalf("Failed to start transaction: %v", err)
	}
	defer tx.Rollback() // Rollback any changes in case the test fails

	user := User{Username: "Test User", Password: "12345678", Email: gofakeit.Email()}

	err = insertUser(tx, &user)
	assert.NoError(t, err, "Failed to insert user 1")
	assert.NotZero(t, user.ID, "User1 ID should be non-zero after insertion")
	fmt.Println(user.ID)
}
