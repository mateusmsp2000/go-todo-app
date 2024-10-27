package services

import (
	"github.com/google/uuid"
	"go-todo-app/commands"
	"go-todo-app/models"
)

type TaskService struct {
	AddTaskCmd    *commands.AddTaskCommand
	UpdateTaskCmd *commands.UpdateTasksCommand
	RemoveTaskCmd *commands.RemoveTasksCommand
}

func NewTaskService(
	addTaskCmd *commands.AddTaskCommand,
	updateTaskCmd *commands.UpdateTasksCommand,
	removeTaskCmd *commands.RemoveTasksCommand) ITaskService {
	return &TaskService{
		AddTaskCmd:    addTaskCmd,
		UpdateTaskCmd: updateTaskCmd,
		RemoveTaskCmd: removeTaskCmd,
	}
}

func (service *TaskService) CompleteInsert(tasks []models.Task) ([]models.Task, error) {
	return service.AddTaskCmd.Execute(tasks)
}

func (service *TaskService) CompleteUpdate(tasks []models.Task) error {
	return service.UpdateTaskCmd.Execute(tasks)
}

func (service *TaskService) CompleteDelete(ids []uuid.UUID) error {
	return service.RemoveTaskCmd.Execute(ids)
}
