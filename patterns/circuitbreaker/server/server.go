package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time = time.Now()

func Server() {
	e := gin.Default()
	e.GET("/ping", func(ctx *gin.Context) {
		// To simulate the upstream service is down, we return a 500 error code to the client in the first 5 seconds on startup.
		if time.Since(startTime) < 5*time.Second {
			ctx.String(http.StatusInternalServerError, "pong")
			return
		}
		ctx.String(http.StatusOK, "pong")
	})

	fmt.Printf("Starting server at port 8080\n")
	e.Run(":8080")
}
