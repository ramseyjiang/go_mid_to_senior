package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var servers = []string{
	"http://localhost:8081",
	"http://localhost:8082",
}

var counter int32

func getServer() string {
	index := atomic.AddInt32(&counter, 1)
	return servers[index%int32(len(servers))]
}

func CreateRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", handleRequest)

	return router
}

func StartServer(router *gin.Engine) {
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func handleRequest(c *gin.Context) {
	server := getServer()
	remote, err := url.Parse(server)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse remote URL"})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(c.Writer, c.Request)
}
