package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/gin_handle_errors/middleware"
)

func main() {
	r := gin.Default()
	// use middleware
	r.Use(middleware.AnyErrorWrapper())

	r.GET("/", handleSome)
	_ = r.Run()
}

func handleSome(c *gin.Context) {
	_ = c.Error(errors.New("What is this"))
	c.String(http.StatusOK, "Hello world!")
}
