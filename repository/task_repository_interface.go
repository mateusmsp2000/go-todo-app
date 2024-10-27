package repository

import (
	"github.com/google/uuid"
	"go-todo-app/models"
)

type ITaskRepository interface {
	AddTasks(tasks []models.Task) error
	ListTasks() ([]models.Task, error)
	UpdateTasks(tasks []models.Task) error
	DeleteTasks(ids []uuid.UUID) error
}
