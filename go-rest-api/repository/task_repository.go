package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAll(tasks *[]model.Task, userID uint) error
	GetByID(task *model.Task, userID uint, taskID uint) error
	Create(task *model.Task) error
	Update(task *model.Task, userID uint, taskID uint) error
	Delete(userID uint, taskID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAll(tasks *[]model.Task, userID uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userID).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetByID(task *model.Task, userID uint, taskID uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userID).First(task, taskID).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Create(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Update(task *model.Task, userID uint, taskID uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", taskID, userID).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) Delete(userID uint, taskID uint) error {
	result := tr.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
