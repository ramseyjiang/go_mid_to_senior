package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	)`

	_, err := db.Exec(query)
	return err
}

func (u *User) CreateUser(db *sql.DB) (int, error) {
	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`

	var id int
	err := db.QueryRow(query, u.Name, u.Age).Scan(&id)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (u *User) GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, name, age FROM users WHERE id = $1`

	row := db.QueryRow(query, id)
	user := &User{}

	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("user not found")
			return nil, fmt.Errorf("user not found")
		}

		log.Error().Err(err).Msg("failed to get user")
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
