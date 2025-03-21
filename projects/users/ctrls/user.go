package ctrls

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/models"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/service"
)

type userCtrlInterface interface {
	Create(*gin.Context)
	GetAll(*gin.Context)
}

type userCtrl struct{}

var UserCtrl userCtrlInterface

func init() {
	UserCtrl = new(userCtrl)
}

func (ctrl *userCtrl) Create(c *gin.Context) {
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

func (ctrl *userCtrl) GetAll(c *gin.Context) {
	resp, _ := service.UserGetAll()

	c.JSON(http.StatusOK, resp)
}
