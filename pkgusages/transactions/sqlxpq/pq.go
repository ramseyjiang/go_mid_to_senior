package sqlxpq

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const createUsersTableQuery = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL
);
`

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

func connectToDatabase() (*sqlx.DB, error) {
	host := "localhost"
	port := 5432
	user := "root"
	pwd := ""
	dbName := "practice"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		pwd,
		host,
		port,
		dbName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createUsersTableQuery)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	return db, nil
}

func insertUser(tx *sqlx.Tx, user *User) error {
	query := `INSERT INTO users (username, email) VALUES (:username, :email) RETURNING id`
	rows, err := tx.NamedQuery(query, user)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID)
		if err != nil {
			return err
		}
	}

	return rows.Err()
}
