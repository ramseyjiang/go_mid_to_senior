package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang_learn/projects/users/config"
	"golang_learn/projects/users/controllers"
	"golang_learn/projects/users/migrations"
)

var (
	defaultPort = "8088"
	path        = ".env"
)

func init() {
	if err := godotenv.Load(path); err != nil {
		panic(err)
	}
}

/**
	In users folder, run.

	% go test ./tests
	ok      golang_learn/projects/users/tests  0.122s

	% go test -v ./tests
	...
	--- PASS: TestSuite (0.11s)
    --- PASS: TestSuite/TestCreateUser (0.01s)
	PASS
	ok      golang_learn/projects/users/tests  0.131s

	% go run server.go
	...
	[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
	 - using env:   export GIN_MODE=release
	 - using code:  gin.SetMode(gin.ReleaseMode)

	[GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
	[GIN-debug] GET    /users                    --> golang_learn/projects/users/controllers.userControllerInterface.GetAll-fm (2 handlers)
	[GIN-debug] POST   /users                    --> golang_learn/projects/users/controllers.userControllerInterface.Create-fm (2 handlers)
	[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
	[GIN-debug] Listening and serving HTTP on :8080
*/
func main() {
	config.ConnectGorm()
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	migrations.MigrateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// Testing Port
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/users", controllers.UserController.GetAll)
	router.POST("/users", controllers.UserController.Create)

	_ = router.Run(":" + port)
}
