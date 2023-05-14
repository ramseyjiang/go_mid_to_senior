package main

import (
	"fmt"
	"go-dock/data"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/animal/:name", func(c *gin.Context) {
		animal, err := data.GetAnimal(c.Param("name"))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, animal)
	})

	r.GET("db/conn", func(c *gin.Context) {
		host := "0.0.0.0"
		port := "5432"
		username := "test"
		pwd := "test"
		dbName := "db"

		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
			host,
			username,
			pwd,
			dbName,
			port,
		)
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}

		// 	sql := `
		// 	CREATE TABLE IF NOT EXISTS users (
		// 		id SERIAL PRIMARY KEY,
		// 		name TEXT NOT NULL
		// 	);
		// `
		//
		// 	result := db.Exec(sql)
		// 	if result.Error != nil {
		// 		fmt.Println(result.Error)
		// 		return
		// 	}
		//
		// 	type User struct {
		// 		ID   int
		// 		Name string
		// 	}
		//
		// 	var user User
		// 	err = db.First(&user, 1).Error
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}
		c.JSON(http.StatusOK, gin.H{
			// "name": user.Name,
			// "id":   user.ID,
			"conn": "connect db success",
			"db":   db,
		})
	})

	_ = r.Run(":3000")
}
