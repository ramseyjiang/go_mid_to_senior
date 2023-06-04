package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create reverse proxies for the user and log services
	userURL, _ := url.Parse("http://localhost:8080")
	userProxy := httputil.NewSingleHostReverseProxy(userURL)

	logURL, _ := url.Parse("http://localhost:8081")
	logProxy := httputil.NewSingleHostReverseProxy(logURL)

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin engine
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	router.POST("/user", func(c *gin.Context) {
		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	router.Any("/log/*path", func(c *gin.Context) {
		logProxy.ServeHTTP(c.Writer, c.Request)
	})

	// Start the server
	router.Run(":8082")
}
