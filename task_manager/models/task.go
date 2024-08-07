package models

import (
	"time"
)

type Task struct {
	ID          string    
	Title       string   
	Description string    
	DueDate     time.Time 
	Status      string   
}