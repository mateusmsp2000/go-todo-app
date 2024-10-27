package validators

import (
	"errors"
	"go-todo-app/models"
)

type ITaskValidator interface {
	ValidateInsert(task *models.Task) error
	ValidateUpdate(task *models.Task) error
	ValidateRemove(task *models.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (v *taskValidator) ValidateInsert(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

func (v *taskValidator) ValidateUpdate(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}
	return nil
}

func (v *taskValidator) ValidateRemove(task *models.Task) error {
	if task.Completed {
		return errors.New("cannot remove a completed task")
	}
	return nil
}
