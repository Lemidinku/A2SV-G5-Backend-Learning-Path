package models

import (
	"errors"
	"time"
)

type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
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