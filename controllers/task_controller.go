package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-todo-app/models"
	"go-todo-app/services"
	"net/http"
	"strings"
)

type TaskController struct {
	TaskService services.ITaskService
}

func NewTaskController(taskService services.ITaskService) ITaskController {
	return &TaskController{
		TaskService: taskService,
	}
}

func (c *TaskController) AddTasks(ctx *gin.Context) {
	var tasks []models.Task
	if err := ctx.ShouldBindJSON(&tasks); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTasks, err := c.TaskService.CompleteInsert(tasks)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdTasks)
}

func (c *TaskController) UpdateTasks(ctx *gin.Context) {
	var tasks []models.Task
	if err := ctx.ShouldBindJSON(&tasks); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.TaskService.CompleteUpdate(tasks); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tasks updated successfully"})
}

func (c *TaskController) RemoveTasks(ctx *gin.Context) {
	idsStr := ctx.Query("ids")
	ids := strings.Split(idsStr, ",")
	var taskIDs []uuid.UUID

	for _, idStr := range ids {
		id, err := uuid.Parse(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		taskIDs = append(taskIDs, id)
	}

	if err := c.TaskService.CompleteDelete(taskIDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove tasks"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tasks removed successfully"})
}
