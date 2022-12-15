package tests

import (
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/models"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/service"
	"github.com/stretchr/testify/assert"
)

func (t *SuiteTest) TestCreateUser() {
	_, err := service.UserCreate(models.User{
		Name:  "First",
		Email: "first@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(models.User{
		Name:  "Second",
		Email: "second@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(models.User{
		Name:  "Third",
		Email: "second@gmail.com",
	})
	assert.Error(t.T(), err) // Duplicate Email Error
}
