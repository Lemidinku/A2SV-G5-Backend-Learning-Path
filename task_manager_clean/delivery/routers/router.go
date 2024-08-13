package routers

import (
	"task_manager_clean/delivery/controllers"
	"task_manager_clean/infrastructure"

	"github.com/gin-gonic/gin"
)

func RunTaskManager(controller controllers.TaskController) {
	router := gin.Default()


	router.POST("/login",controller.Login )
	router.POST("/register",controller.Register )

	auth := router.Group("")
	{	
		router.Use(infrastructure.AuthMiddleware())
		auth.GET("/tasks/", controller.GetTasks);
		auth.GET("/tasks/:id", controller.GetTask);
		auth.Group(""); {
			auth.Use(infrastructure.OnlyAdmin())
			auth.DELETE("/tasks/:id", controller.RemoveTask)
			auth.PUT("/tasks/:id", controller.UpdateTask)
			auth.POST("/tasks", controller.AddTask)
			auth.PUT("/promote/:username", controller.PromoteUser)
	};
	}
	
	router.Run()
}