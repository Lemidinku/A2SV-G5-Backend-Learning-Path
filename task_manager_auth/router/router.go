package router

import (
	"task_manager_auth/controllers"
	"task_manager_auth/middleware"

	"github.com/gin-gonic/gin"
)

func RunTaskManager() {
	router := gin.Default()

	taskControllers := controllers.NewTaskController()

	router.POST("/login",taskControllers.Login )
	router.POST("/register",taskControllers.Register )
	auth := router.Group("")
	{	
		router.Use(middleware.AuthMiddleware())
		auth.GET("/tasks/", taskControllers.GetTasks);
		auth.GET("/tasks/:id", taskControllers.GetTask);
		auth.Group(""); {
			auth.Use(middleware.OnlyAdmin())
			auth.DELETE("/tasks/:id", taskControllers.RemoveTask)
			auth.PUT("/tasks/:id", taskControllers.UpdateTask)
			auth.POST("/tasks", taskControllers.AddTask)
			auth.PUT("/promote/:username", taskControllers.PromoteUser)
	};
	}
	
	router.Run()
}