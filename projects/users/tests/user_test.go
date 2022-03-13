package tests

import (
	"github.com/stretchr/testify/assert"
	"golang_learn/projects/users/entity"
	"golang_learn/projects/users/service"
)

func (t *SuiteTest) TestCreateUser() {
	_, err := service.UserCreate(entity.User{
		Name:  "First",
		Email: "first@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(entity.User{
		Name:  "Second",
		Email: "second@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(entity.User{
		Name:  "Third",
		Email: "second@gmail.com",
	})
	assert.Error(t.T(), err) // Duplicate Email Error
}
