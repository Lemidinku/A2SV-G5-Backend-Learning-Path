package models


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
