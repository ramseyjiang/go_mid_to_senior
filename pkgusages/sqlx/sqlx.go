package sqlxpkg

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int    `data:"id"`
	FirstName string `data:"first_name"`
	LastName  string `data:"last_name"`
	Email     string `data:"email"`
	Mobile    string `data:"mobile"`
}

// It is used only when a table does not exist.
// var schema = "CREATE TABLE `users` (" +
// 	"`id` varchar(100) NOT NULL," +
// 	"`first_name` varchar(255) NOT NULL," +
// 	"`last_name` varchar(255) NOT NULL," +
// 	"`email` varchar(255) NOT NULL," +
// 	"`mobile` varchar(255) NOT NULL," +
// 	"PRIMARY KEY (`id`)," +
// 	"UNIQUE KEY `email` (`email`)," +
// 	"UNIQUE KEY `mobile` (`mobile`))"

func Trigger() {
	db, err := sqlx.Connect("mysql", "root:12345678@(localhost:3306)/go_web")
	if err != nil {
		log.Fatalln(err)
	}

	// Only exec when the table needs to create. If the table has been created, it will have an error.
	// db.MustExec(schema)

	id := insertUser(db)
	fmt.Printf("Created user with id:%d\n", id)

	getUserByID(db, id)
	updateUserByID(db, id)
	delUserByID(db, id-1)

	users, err := getAllUsers(db)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(users)
}

func insertUser(db *sqlx.DB) int64 {
	res, err := db.Exec("INSERT INTO users (id, first_name, last_name, email, mobile) VALUES(\"3153715d-35a1-4bb3-99ea-3f7d371dd911\", \"Peter\", \"DogFather\", \"davy@gmail.com\", 4732819)")
	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}

func getUserByID(db *sqlx.DB, id int64) {
	var users User
	err := db.Get(&users, "select * from users where id=?", id)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}

func updateUserByID(db *sqlx.DB, id int64) {
	_, err := db.Exec("UPDATE users set email=\"Test@gmail.com\" where id=?", id)
	if err != nil {
		panic(err)
	}
}

func delUserByID(db *sqlx.DB, id int64) {
	_, err := db.Exec("DELETE FROM users where id=?", id)
	if err != nil {
		panic(err)
	}
}

func getAllUsers(db *sqlx.DB) ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users")

	return users, err
}
