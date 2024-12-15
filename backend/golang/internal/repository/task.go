package repository

import (
	"myapp/infrastructure/entity"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task entity.Task) error
	GetTask(taskID string) (entity.Task, error)
	UpdateTask(task entity.Task) error
	DeleteTask(taskID string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) CreateTask(task entity.Task) error {
	if err := r.db.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetTask(taskID string) (entity.Task, error) {
	var task entity.Task
	if err := r.db.Where("task_id = ?", taskID).First(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *taskRepository) UpdateTask(task entity.Task) error {
	if err := r.db.Save(&task).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTask(taskID string) error {
	var task entity.Task
	if err := r.db.Where("task_id = ?", taskID).Delete(&task).Error; err != nil {
		return err
	}
	return nil
}
