package datadog

import (
	"context"
	"database/sql"
	"log"
)

type UserDBRepository struct {
	connection *sql.DB
}

func NewUserDBRepository(connection *sql.DB) *UserDBRepository {
	return &UserDBRepository{
		connection: connection,
	}
}

type User struct {
	lastname string
}

func (r *UserDBRepository) FilterByLastname(ctx context.Context, lastname string) ([]User, error) {
	var users []User

	rows, _ := r.connection.Query("SELECT * FROM users WHERE lastname = ?", lastname)
	log.Println(rows)
	return users, nil
}
