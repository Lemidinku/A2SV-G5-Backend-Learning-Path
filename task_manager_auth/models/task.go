package models

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