package models

import "errors"

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
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}
