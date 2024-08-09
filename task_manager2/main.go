package main

import (
	"task_manager2/router"
	"task_manager2/database"
)

func main() {

		database.ConnectMongoDB()
		router.RunTaskManager()

		
}
