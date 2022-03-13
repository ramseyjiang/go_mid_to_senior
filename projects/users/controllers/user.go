package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang_learn/projects/users/entity"
	"golang_learn/projects/users/service"
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
	var input entity.User
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
