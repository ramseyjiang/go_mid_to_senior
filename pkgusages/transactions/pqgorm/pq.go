package pqgorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
}

func connectToDatabase() (*gorm.DB, error) {
	host := "localhost"
	port := 5432
	user := "root"
	pwd := ""
	dbName := "practice"

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbName,
		pwd,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Migrate the schema
	err = db.AutoMigrate(&User{})

	return db, err
}
