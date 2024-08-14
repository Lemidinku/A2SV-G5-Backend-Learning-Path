package controllers

import (
	"net/http"
	"task_manager_clean/domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
	userUsecase domain.UserUsecase
}

func NewTaskController(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) *TaskController {
	return &TaskController{taskUsecase: taskUsecase,userUsecase: userUsecase}
}

func (taskController *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := taskController.taskUsecase.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func (taskController *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := taskController.taskUsecase.GetTask(id)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"task": task})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func (taskController *TaskController) AddTask(ctx *gin.Context) {
	var newTask domain.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedTask, err := taskController.taskUsecase.AddTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"task": addedTask})
}

func (taskController *TaskController) RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")

	err := taskController.taskUsecase.RemoveTask(id)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func (taskController *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask domain.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := taskController.taskUsecase.UpdateTask(id, updatedTask)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{"task": updatedTask})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// register a new user
func (taskController *TaskController) Register(ctx *gin.Context) {
	var newUser domain.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := taskController.userUsecase.RegisterUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// login a user
func (taskController *TaskController) Login(ctx *gin.Context) {
	var user domain.User
	
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := taskController.userUsecase.GetUser(user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User logged in successfully", "token": token})
}

// promote a user to admin
func (taskController *TaskController) PromoteUser(ctx *gin.Context) {
	username := ctx.Param("username")

	message, err := taskController.userUsecase.PromoteUser(username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
