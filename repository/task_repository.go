package repository

import (
	"github.com/google/uuid"
	"go-todo-app/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) AddTasks(tasks []models.Task) error {
	return r.DB.Create(&tasks).Error
}

func (r *TaskRepository) ListTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) UpdateTasks(tasks []models.Task) error {
	for _, task := range tasks {
		if err := r.DB.Model(&task).Updates(task).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *TaskRepository) DeleteTasks(ids []uuid.UUID) error {
	if err := r.DB.Where("id IN ?", ids).Delete(&models.Task{}).Error; err != nil {
		return err
	}
	return nil
}
