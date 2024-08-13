package domain

import (
	"time"
)

type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
	Status      string    `json:"status" bson:"status"`
}


type Role string
// define role types
const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     Role   `json:"role" bson:"role"`
}


// define the interface for the repository
type TaskRepository interface {
	GetTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	AddTask(task Task) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	RemoveTask(id string) error
}

type UserRepository interface {
	RegisterUser(user User) (User, error)
	GetUser(user User) (string, error)
	PromoteUser(username string) (string, error)
}


// define usecase interface
type TaskUsecase interface {
	GetTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	AddTask(task Task) (Task, error)
	UpdateTask(id string, task Task) (Task, error)
	RemoveTask(id string) error
}

type UserUsecase interface {
	RegisterUser(user User) (User, error)
	GetUser(user User) (string, error)
	PromoteUser(username string) (string, error)
}