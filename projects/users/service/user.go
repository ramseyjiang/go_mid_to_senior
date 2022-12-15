package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/config"
	"github.com/ramseyjiang/go_mid_to_senior/projects/users/models"
	"gorm.io/gorm"
)

func UserCreate(input models.User) (*models.User, error) {
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

func UserGetAll() ([]*models.User, error) {
	db := config.GetDB()

	var users []*models.User
	if err := db.Model(users).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UserGetByEmail(email string) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.Model(user).Where("email like ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
