package router

import (
	"task_manager2/controllers"

	"github.com/gin-gonic/gin"
)

func RunTaskManager() {
	router := gin.Default()

	taskControllers := controllers.NewTaskController()

	router.GET("/tasks", taskControllers.GetTasks)
	router.GET("/tasks/:id", taskControllers.GetTask)
	router.DELETE("/tasks/:id", taskControllers.RemoveTask)
	router.PUT("/tasks/:id", taskControllers.UpdateTask)
	router.POST("/tasks", taskControllers.AddTask)
	
	router.Run()
}