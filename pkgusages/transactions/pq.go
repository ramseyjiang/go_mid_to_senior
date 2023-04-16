package transactions

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func connectToDatabase() (*sqlx.DB, error) {
	host := "localhost"
	port := 5432
	user := "root"
	pwd := ""
	dbName := "pro"
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

	return db, nil
}

func insertUser(tx *sqlx.Tx, user *User) error {
	query := `INSERT INTO users (username, password, email) VALUES (:username, :password, :email) RETURNING id`
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
