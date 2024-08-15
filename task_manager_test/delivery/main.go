package main

import (
	"task_manager_test/delivery/controllers"
	"task_manager_test/delivery/routers"
	"task_manager_test/infrastructure"
	"task_manager_test/repositories"
	"task_manager_test/usecases"

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

		// Initialize the password service
		passwordService := infrastructure.NewPasswordService()
		jwtService := infrastructure.NewJWTService("")

		// Initialize the usecases
		userUsecase := usecases.NewUserUsecase(userRepository, passwordService,jwtService)

		// Initialize the controllers
		controller := controllers.NewTaskController(taskUsecase, userUsecase)
		routers.RunTaskManager(*controller)

}