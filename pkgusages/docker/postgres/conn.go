package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
}

func conn() {
	host := "localhost"
	port := "5432"
	username := "root"
	pwd := "test"
	dbName := "postgres"

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		username,
		dbName,
		pwd,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	sql := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
	`

	result := db.Exec(sql)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	fmt.Println("Table created successfully")

	var user User
	err = db.First(&user, 1).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Name:", user.Name)
}
