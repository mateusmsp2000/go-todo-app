package controllers_reads

import "github.com/gin-gonic/gin"

type ITaskReadController interface {
	ListTasks(ctx *gin.Context)
}
