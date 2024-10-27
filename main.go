package main

import (
	"github.com/gin-gonic/gin"
	"go-todo-app/commands"
	"go-todo-app/controllers"
	controllers_reads "go-todo-app/controllers/reads"
	"go-todo-app/models"
	"go-todo-app/repository"
	"go-todo-app/services"
	"go-todo-app/services/validators"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Repositories
	taskRepo := repository.NewTaskRepository(db)

	// Services Validatos
	validatorTaskService := validators.NewTaskValidator()

	// Commands
	addTaskCmd := commands.NewAddTaskCommand(taskRepo, validatorTaskService)
	updateTaskCmd := commands.NewUpdateTasksCommand(taskRepo, validatorTaskService)
	removeTaskCmd := commands.NewRemoveTasksCommand(taskRepo)

	// Services
	taskService := services.NewTaskService(addTaskCmd, updateTaskCmd, removeTaskCmd)

	// Controllers
	taskController := controllers.NewTaskController(taskService)
	taskReadController := controllers_reads.NewTaskReadController(taskRepo)

	r := gin.Default()

	r.POST("/tasks", taskController.AddTasks)
	r.PUT("/tasks/:id/complete", taskController.UpdateTasks)
	r.DELETE("/tasks", taskController.RemoveTasks)

	//Reads
	r.GET("/tasks", taskReadController.ListTasks)

	r.Run()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
