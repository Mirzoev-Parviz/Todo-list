package repository

import (
	"test/config"
	"test/model"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	user.IsActive = true
	if err := config.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (r *AuthPostgres) GetUser(login string, password string) (model.User, error) {
	var user model.User
	err := config.DB.Where("login = ? AND password = ? AND is_active = TRUE", login, password).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (r *AuthPostgres) Check(login string) (bool, error) {
	var user model.User
	err := config.DB.Where("login = ? AND is_active = TRUE", login).First(&user).Error
	if err != nil {
		return false, err
	}

	if login == user.Login {
		return true, nil
	}

	return false, nil
}
