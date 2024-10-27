package commands

import (
	"go-todo-app/models"
	"go-todo-app/repository"
	"go-todo-app/services/validators"
)

type UpdateTasksCommand struct {
	Repo      repository.ITaskRepository
	Validator validators.ITaskValidator
}

func NewUpdateTasksCommand(repo repository.ITaskRepository, validator validators.ITaskValidator) *UpdateTasksCommand {
	return &UpdateTasksCommand{Repo: repo, Validator: validator}
}

func (c *UpdateTasksCommand) Execute(tasks []models.Task) error {
	for _, task := range tasks {
		if err := c.Validator.ValidateUpdate(&task); err != nil {
			return err
		}
	}

	if err := c.Repo.UpdateTasks(tasks); err != nil {
		return err
	}

	return nil
}
