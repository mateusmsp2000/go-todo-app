package controllers_reads

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/repository"
	"net/http"
)

type TaskReadController struct {
	TaskRepo repository.ITaskRepository
}

func NewTaskReadController(taskRepo repository.ITaskRepository) ITaskReadController {
	return &TaskReadController{
		TaskRepo: taskRepo,
	}
}

func (c *TaskReadController) ListTasks(ctx *gin.Context) {
	tasks, err := c.TaskRepo.ListTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list tasks"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
