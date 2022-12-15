package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/config"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/ctrls"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/migrations"
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
	ok      github.com/ramseyjiang/go_mid_to_senior/projects/users/tests  0.122s

	% go test -v ./tests
	...
	--- PASS: TestSuite (0.11s)
    --- PASS: TestSuite/TestCreateUser (0.01s)
	PASS
	ok      github.com/ramseyjiang/go_mid_to_senior/projects/users/tests  0.131s

	% go run server.go
	...
	[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
	 - using env:   export GIN_MODE=release
	 - using code:  gin.SetMode(gin.ReleaseMode)

	[GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
	[GIN-debug] GET    /users                    --> github.com/ramseyjiang/go_mid_to_senior/projects/users/ctrls.userCtrlInterface.GetAll-fm (2 handlers)
	[GIN-debug] POST   /users                    --> github.com/ramseyjiang/go_mid_to_senior/projects/users/ctrls.userCtrlInterface.Create-fm (2 handlers)
	[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
	[GIN-debug] Listening and serving HTTP on :8080
*/
func main() {
	config.ConnectGorm()
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println("DB connected error:", err.Error())
		}
	}(sqlDB)

	migrations.MigrateTable()
	port := os.Getenv("SERVER_PORT")

	router := gin.New()
	router.Use(gin.Recovery())

	// Testing Port
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/users", ctrls.UserCtrl.GetAll)
	router.POST("/users", ctrls.UserCtrl.Create)

	_ = router.Run(":" + port)
}
