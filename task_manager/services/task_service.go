package services

import (
	"errors"
	. "task_manager/models"
	"time"

	"github.com/google/uuid"
)


type TaskService struct {
	Tasks []Task
}


var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "Description for Task 1", DueDate: time.Now(), Status: "In Progress"},
	{ID: "2", Title: "Task 2", Description: "Description for Task 2", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
}

func NewTaskService() *TaskService {
	task_service := TaskService{Tasks: tasks}
	return &task_service
}



func (task_service *TaskService) GetTasks() []Task {
	return task_service.Tasks
}


func (task_service *TaskService) GetTask(id string) (Task, error)  {
	for _, task := range task_service.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, errors.New("task not found")
}

func (task_service *TaskService) AddTask(newTask Task) Task {
	// use random string as an ID
	newTask.ID = uuid.New().String()
	task_service.Tasks = append(task_service.Tasks, newTask)

	return newTask

}

func (task_service *TaskService) UpdateTask(id string, updatedTask Task) (Task,error) {
	for i, task := range task_service.Tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				task_service.Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				task_service.Tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				task_service.Tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				task_service.Tasks[i].Status = updatedTask.Status
			}
			
			return task_service.Tasks[i], nil
		}
	}
	return Task{},errors.New("task not found")
}


func (task_service *TaskService) RemoveTask(id string) error {
	for i, val := range task_service.Tasks {
		if val.ID == id {
			task_service.Tasks = append(task_service.Tasks[:i], task_service.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}


