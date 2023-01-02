package api

import (
	"fmt"

	ginopentracing "github.com/Bose/go-gin-opentracing"
	"github.com/gin-gonic/gin"
)

func path(endpoint string) string {
	return fmt.Sprintf("/api/%s", endpoint)
}

type Router struct {
	router *gin.Engine
}

func (api *Router) Serve(addr string) error {
	return api.router.Run(addr)
}

func NewRouter() *Router {
	r := gin.Default()
	p := ginopentracing.OpenTracer([]byte("api-request-"))
	// tell gin to use the middleware
	r.Use(p)

	apiRouter := &Router{
		r,
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Start tracing")
	})
	r.GET(path("parent"), SpanFromParent)
	r.GET(path("func"), SpanHasFuncName)
	r.GET("header", SpanFromHeader)
	_ = r.Run(":9000")

	return apiRouter
}
