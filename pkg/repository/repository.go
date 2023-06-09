package repository

import (
	"test/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string, password string) (model.User, error)
	Check(login string) (bool, error)
}

type TODO interface {
	CreateTODO(userId int, todo model.TODO) (int, error)
	GetTODOS(userId int) ([]model.TODO, error)
	UpdateTODO(todo model.TODO, id, userId int) error
	DeleteTODO(id, userId int) error
}

type Repository struct {
	Authorization
	TODO
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TODO:          NewTodoPostgres(db),
	}
}
