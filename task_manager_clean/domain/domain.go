package domain

import (
	"errors"
	"time"
)

type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title" validate:"required"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"dueDate" bson:"dueDate"`
	Status      TaskStatus    `json:"status" bson:"status"`
}
type TaskStatus string
// define role types
const (
	Completed TaskStatus = "completed"
	Pending  TaskStatus = "pending"
)

func (task *Task) ValidateTask() error {
	if (task.Title == "") {
		return errors.New("title is required")
	}
	if (len(task.Title) < 3) {
		return errors.New("title must be at least 3 characters")
	}
	if (task.DueDate.IsZero()) {
		return errors.New("due date is required")
	}
	if (task.Status == "") {
		return errors.New("status is required")
	}
	if (task.Status != Completed && task.Status != Pending) {
		return errors.New("status must be either completed or pending")
	}
	return nil
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

func (user *User) ValidateUser() error {
	if (user.Username == "") {
		return errors.New("username is required")
	}
	if (user.Password == "") {
		return errors.New("password is required")
	}
	if (len(user.Username) < 3) {
		return errors.New("username must be at least 3 characters")
	}
	if (len(user.Password) < 6) {
		return errors.New("password must be at least 6 characters")
	}
	return nil
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
	GetUser(user User) (User, error)
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