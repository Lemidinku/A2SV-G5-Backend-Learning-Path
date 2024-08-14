package main

import (
	"task_manager_clean/delivery/controllers"
	"task_manager_clean/delivery/routers"
	"task_manager_clean/infrastructure"
	"task_manager_clean/repositories"
	"task_manager_clean/usecases"

	"log"

	"github.com/joho/godotenv"
)


func main() {

	    // Load the .env file during initialization
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		// Connect to the database
		client, err := infrastructure.ConnectToMongoDB()
		if err != nil {
			log.Fatal(err)
		}

		// Create the collections
		TaskCollection, _ := infrastructure.CreateCollection(client, "task_manager", "tasks")
		UserCollection,_ := infrastructure.CreateCollection(client, "task_manager", "users")

		// Initialize the repositories
		taskRepository := repositories.NewTaskRepository(TaskCollection)
		userRepository := repositories.NewUserRepository(UserCollection)

		// Initialize the usecases
		var taskUsecase  = usecases.NewTaskUsecase(taskRepository)
		userUsecase := usecases.NewUserUsecase(userRepository)

		// Initialize the controllers
		controller := controllers.NewTaskController(taskUsecase, userUsecase)
		routers.RunTaskManager(*controller)

}