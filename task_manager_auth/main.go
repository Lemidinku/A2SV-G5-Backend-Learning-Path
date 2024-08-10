package main

import (
	"task_manager_auth/router"
	"task_manager_auth/database"
	"log"
	"github.com/joho/godotenv"
)

func main() {
	    // Load the .env file during initialization
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
		database.ConnectMongoDB()
		router.RunTaskManager()
}
