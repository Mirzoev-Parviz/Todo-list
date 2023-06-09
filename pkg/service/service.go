package service

import (
	"test/model"
	"test/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
	IsUserExist(login string) (bool, error)
}

type TODO interface {
	CreateTODO(userId int, todo model.TODO) (int, error)
	GetTODOS(userID int) ([]model.TODO, error)
	UpdateTODO(todo model.TODO, id, userId int) error
	DeleteTODO(id, userId int) error
}

type Service struct {
	Authorization
	TODO
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TODO:          NewTODOService(repos.TODO),
	}
}
