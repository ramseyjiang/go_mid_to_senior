package reverseproxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Define the routes for the API Gateway
	router.Any("/service1/*path", createReverseProxy("http://localhost:8081"))
	router.Any("/service2/*path", createReverseProxy("http://localhost:8082"))

	return router
}

func startServer(router *gin.Engine) {
	// Start the API Gateway
	router.Run(":8080")
}

func createReverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL
		targetURL, _ := url.Parse(target)

		// Create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Modify the request
		c.Request.URL.Scheme = targetURL.Scheme
		c.Request.URL.Host = targetURL.Host
		c.Request.URL.Path = c.Param("path")

		// Let the reverse proxy do its job
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
