package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ramseyjiang/go_mid_to_senior/pkgusages/stock/controllers"
)

func main() {
	e := echo.New()            // Middleware
	e.Use(middleware.Logger()) // Logger
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Root route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.POST("/price", controllers.GrabPrice) // Price endpoint
	// Server
	e.Logger.Fatal(e.Start(":8001"))
}
