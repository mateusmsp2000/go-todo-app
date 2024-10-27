package services

import (
	"github.com/google/uuid"
	"go-todo-app/models"
)

type ITaskService interface {
	CompleteInsert(tasks []models.Task) ([]models.Task, error)
	CompleteUpdate(tasks []models.Task) error
	CompleteDelete(ids []uuid.UUID) error
}
