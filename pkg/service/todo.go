package service

import (
	"test/model"
	"test/pkg/repository"
)

type TODOService struct {
	repo repository.TODO
}

func NewTODOService(repo repository.TODO) *TODOService {
	return &TODOService{repo: repo}
}

func (t *TODOService) CreateTODO(userId int, todo model.TODO) (int, error) {
	return t.repo.CreateTODO(userId, todo)
}
func (t *TODOService) GetTODOS(userId int) ([]model.TODO, error) {
	return t.repo.GetTODOS(userId)
}
func (t *TODOService) UpdateTODO(todo model.TODO, id, userId int) error {
	return t.repo.UpdateTODO(todo, id, userId)
}
func (t *TODOService) DeleteTODO(id, userId int) error {
	return t.repo.DeleteTODO(id, userId)
}
