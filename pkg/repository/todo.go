package repository

import (
	"test/config"
	"test/model"

	"gorm.io/gorm"
)

type TodoPostgres struct {
	db *gorm.DB
}

func NewTodoPostgres(db *gorm.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}
func (t *TodoPostgres) CreateTODO(userId int, todo model.TODO) (int, error) {
	todo.UserID = userId
	todo.IsActive = true
	if err := config.DB.Create(&todo).Error; err != nil {
		return 0, err
	}

	return todo.Id, nil
}
func (t *TodoPostgres) GetTODOS(userId int) (todoList []model.TODO, err error) {
	if err = config.DB.Where("user_id = ? AND is_active = TRUE", userId).Find(&todoList).Error; err != nil {
		return []model.TODO{}, err
	}

	return todoList, nil
}
func (t *TodoPostgres) UpdateTODO(todo model.TODO, id, userId int) error {
	if err := config.DB.Where("id = ? AND user_id = ? AND is_active = TRUE").Updates(&todo).Error; err != nil {
		return err
	}

	return nil
}
func (t *TodoPostgres) DeleteTODO(id, userId int) error {
	var todo model.TODO
	if err := config.DB.Where("id = ? AND user_id = ?", id, userId).First(&todo).Error; err != nil {
		return err
	}
	todo.IsActive = false

	if err := config.DB.Save(&todo).Error; err != nil {
		return err
	}

	return nil
}
