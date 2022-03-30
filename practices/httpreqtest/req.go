package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Result struct {
	Value int `json:"value"`
}

type AppHandler interface {
	Sum(x int, y int) Result
	Multiply(x int, y int) Result
}

type AppHandlerStruct struct{}

func (a *AppHandlerStruct) Sum(x int, y int) (r Result) {
	r.Value = Sum(x, y)
	return
}
func (a *AppHandlerStruct) Multiply(x int, y int) (r Result) {
	r.Value = Multiply(x, y)
	return
}

func Sum(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

// MultiplyRequest Request body payload of the 'POST /multiply' endpoint
type MultiplyRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CreateApp(handler AppHandler) *fiber.App {
	// Create new instance
	app := fiber.New()
	app.Get("/sum", sum(handler))
	app.Post("/multiply", multiply(handler))

	return app
}

/*
	Use a postman to send a json data request as below.
	Example request body payload:
	{
		"x": 5,
		"y": 4
	}
*/
func multiply(handler AppHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Try to parse request body to 'req' object
		req := new(MultiplyRequest)
		if err := c.BodyParser(req); err != nil {
			log.Println(err)

			// Return '400 Bad Request' with a text message
			return c.Status(http.StatusBadRequest).SendString("Invalid payload")
		}

		// Use "Multiply" handler to calculate the multiplication of both values
		r := handler.Multiply(req.X, req.Y)

		// Return '200 OK' with a 'Result' object containing the calculated value
		return c.Status(http.StatusOK).JSON(r)
	}
}

// Example URL: "/sum?x=5&y=4"
func sum(handler AppHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get values from query parameters
		x, _ := strconv.Atoi(c.Query("x"))
		y, _ := strconv.Atoi(c.Query("y"))

		// Use "Sum" handler to calculate the sum of both values
		r := handler.Sum(x, y)

		// Return '200 OK' with a 'Result' object containing the calculated value
		return c.Status(http.StatusOK).JSON(r)
	}
}

func main() {
	// Create the App
	app := CreateApp(&AppHandlerStruct{})

	// Listen to Port 3000
	err := app.Listen(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
