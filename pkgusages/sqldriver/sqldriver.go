package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id        int
	name      string
	email     string
	createdAt time.Time
}

// Create a new table
func createUsersTable(db *sql.DB) {
	query := `
		    CREATE TABLE users (
		        id INT AUTO_INCREMENT,
		        name varchar(255) NOT NULL,
		        email varchar(255) NOT NULL,
		        created_at DATETIME DEFAULT NULL,
		        PRIMARY KEY (id)
		    );`

	if _, err := db.Exec(query); err != nil {
		// If the users table has been created, go ahead, don't show the error.
		if err.Error() != "Error 1050: Table 'users' already exists" {
			log.Fatal(err)
		}
	}
}

// Insert a new user
func insertRow(db *sql.DB) (int64, error) {
	name := "john doe"
	email := "john@gmail.com"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)`, name, email, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

// Query a single user
func getRowByID(db *sql.DB, userID int64) {
	var (
		id        int
		name      string
		email     string
		createdAt time.Time
	)

	query := "SELECT id, name, email, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, userID).Scan(&id, &name, &email, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, name, email, createdAt)
}

// Get all users
func getAllRows(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, email, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.name, &u.email, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users)
}

// Delete user by userID
func delRowByID(db *sql.DB, userID int64) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, userID)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete all users
func delAllRows(db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// rules is db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	db, err := sql.Open("mysql", "root:12345678@(127.0.0.1:3306)/go_web?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	if err1 := db.Ping(); err1 != nil {
		log.Fatal(err1)
	}

	defer db.Close()

	createUsersTable(db)
	uID, _ := insertRow(db)
	fmt.Printf("Last insert user id is %v \n", uID)

	getRowByID(db, uID)
	// It will be used when delete row by ID
	delRowByID(db, uID)

	// It will be used when delete all rows
	delAllRows(db)
	getAllRows(db)
}
