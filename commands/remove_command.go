package commands

import (
	"github.com/google/uuid"
	"go-todo-app/repository"
)

type RemoveTasksCommand struct {
	Repo repository.ITaskRepository
}

func NewRemoveTasksCommand(repo repository.ITaskRepository) *RemoveTasksCommand {
	return &RemoveTasksCommand{Repo: repo}
}

func (r *RemoveTasksCommand) Execute(ids []uuid.UUID) error {
	if err := r.Repo.DeleteTasks(ids); err != nil {
		return err
	}
	return nil
}
