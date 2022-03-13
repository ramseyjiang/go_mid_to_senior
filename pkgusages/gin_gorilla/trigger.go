package gingorilla

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"golang_learn/pkgusages/gin_gorilla/ws"
)

func Trigger() {
	go ws.Manager.Start()

	r := gin.Default()
	r.GET("/json", ws.JSONAPI)
	r.GET("/text", ws.TextAPI)
	r.GET("/chat", ws.ChatHandler)
	r.GET("/pong", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// static files
	r.Use(static.Serve("/", static.LocalFile("./", true)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./index.html")
	})

	_ = r.Run(":8000")
}
