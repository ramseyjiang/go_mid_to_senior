package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"golang_learn/projects/users/config"
	"golang_learn/projects/users/entity"
	"gorm.io/gorm"
)

func UserCreate(input entity.User) (*entity.User, error) {
	input.Email = strings.ToLower(input.Email)
	_, err := UserGetByEmail(input.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err == nil {
		return nil, errors.New("email not available")
	}

	db := config.GetDB()
	input.ID = uuid.New().String()
	if err1 := db.Model(input).Create(&input).Error; err1 != nil {
		return nil, err1
	}

	return &input, nil
}

func UserGetAll() ([]*entity.User, error) {
	db := config.GetDB()

	var users []*entity.User
	if err := db.Model(users).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UserGetByEmail(email string) (*entity.User, error) {
	db := config.GetDB()

	var user entity.User
	if err := db.Model(user).Where("email like ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
