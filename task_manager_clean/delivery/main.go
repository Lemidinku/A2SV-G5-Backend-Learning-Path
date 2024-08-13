package main

import (
	"task_manager_clean/delivery/controllers"
	"task_manager_clean/delivery/routers"
	"task_manager_clean/repositories"
	"task_manager_clean/usecases"

	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main() {

	    // Load the .env file during initialization
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	
		TaskCollection := client.Database("task_manager").Collection("tasks")
		UserCollection := client.Database("task_manager").Collection("users")

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