package controllers

import (
	"net/http"
	. "task_manager/models"
	"task_manager/services"

	"github.com/gin-gonic/gin"
)


type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: services.NewTaskService(),
	}
}

func (taskController *TaskController) GetTasks(ctx *gin.Context) {
	tasks := taskController.taskService.GetTasks()
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func (taskController *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := taskController.taskService.GetTask(id)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"task": task})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}


func (taskController *TaskController) AddTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	addedTask := taskController.taskService.AddTask(newTask)
	ctx.JSON(http.StatusCreated, gin.H{"task": addedTask})
}


func (taskController *TaskController) RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")
	
	err := taskController.taskService.RemoveTask(id)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})

}


func (taskController *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := taskController.taskService.UpdateTask(id, updatedTask)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"task": updatedTask})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}


