package main

import (
	"go-dock/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/animal/:name", func(c *gin.Context) {
		animal, err := db.GetAnimal(c.Param("name"))
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusOK, animal)
	})

	_ = r.Run(":3000")
}
