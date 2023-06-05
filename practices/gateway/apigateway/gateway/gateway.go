package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

const userURL = "http://localhost:8080"
const logURL = "http://localhost:8081"

func main() {
	router := createRouter()
	router.Run(":8082")
}

func createRouter() *gin.Engine {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Define the routes for the API Gateway
	router.GET("/user/list", createReverseProxy(userURL, "/user/list"))
	router.POST("/user/create", createReverseProxy(userURL, "/user/create"))
	router.GET("/log/add", createReverseProxy(logURL, "/log/add"))

	return router
}

func createReverseProxy(target string, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL
		targetURL, _ := url.Parse(target)

		// Create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Modify the request
		c.Request.URL.Scheme = targetURL.Scheme
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Path = path

		// Let the reverse proxy do its job
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
