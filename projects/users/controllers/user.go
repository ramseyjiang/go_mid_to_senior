package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/models"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/service"
)

type userControllerInterface interface {
	Create(*gin.Context)
	GetAll(*gin.Context)
}

type userController struct{}

var UserController userControllerInterface

func init() {
	UserController = new(userController)
}

func (controller *userController) Create(c *gin.Context) {
	var input models.User
	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	resp, err := service.UserCreate(input)
	if err != errors.New("email not available") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email Not Available",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (controller *userController) GetAll(c *gin.Context) {
	resp, _ := service.UserGetAll()

	c.JSON(http.StatusOK, resp)
}
