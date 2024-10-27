package controllers

import "github.com/gin-gonic/gin"

type ITaskController interface {
	AddTasks(ctx *gin.Context)
	UpdateTasks(ctx *gin.Context)
	RemoveTasks(ctx *gin.Context)
}
