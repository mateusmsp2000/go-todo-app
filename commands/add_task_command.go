package commands

import (
	"go-todo-app/models"
	"go-todo-app/repository"
	"go-todo-app/services/validators"
)

type AddTaskCommand struct {
	Repo      repository.ITaskRepository
	Validator validators.ITaskValidator
}

func NewAddTaskCommand(repo repository.ITaskRepository, validator validators.ITaskValidator) *AddTaskCommand {
	return &AddTaskCommand{Repo: repo, Validator: validator}
}

func (c *AddTaskCommand) Execute(tasks []models.Task) ([]models.Task, error) {
	for _, task := range tasks {
		if err := c.Validator.ValidateInsert(&task); err != nil {
			return nil, err
		}
	}

	if err := c.Repo.AddTasks(tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
