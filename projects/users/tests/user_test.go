package tests

import (
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/models"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/service"
	"github.com/stretchr/testify/assert"
)

func (t *SuiteTest) TestCreateUser() {
	_, err := service.UserCreate(models.User{
		FirstName: "First",
		LastName:  "test",
		Email:     "first@gmail.com",
		Mobile:    022012311,
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(models.User{
		FirstName: "Second",
		LastName:  "test",
		Email:     "second@gmail.com",
		Mobile:    022012312,
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(models.User{
		FirstName: "Third",
		LastName:  "test",
		Email:     "second@gmail.com",
		Mobile:    022012313,
	})
	assert.Error(t.T(), err) // Duplicate Email Error
}
