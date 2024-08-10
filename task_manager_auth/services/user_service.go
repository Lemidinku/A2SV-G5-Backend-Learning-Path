package services

import (
	"context"
	"errors"
	"task_manager_auth/database"
	"task_manager_auth/models"

	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
	
}

func NewUserService() *UserService {
	return &UserService{
		collection: database.UserCollection,
	}
}


// register a new user
func (service *UserService) RegisterUser(newUser models.User) (models.User, error) {

	// check if user_name already exists
	userNameCount ,err := service.collection.CountDocuments(context.Background(), bson.M{"username": newUser.Username})
	if err != nil {
		return models.User{}, err
	}
	if userNameCount > 0 {
		return models.User{}, errors.New("username already taken")
	}
	newUser.ID = uuid.NewString()

	// check if user is the first user to be registered
	count, err := service.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return models.User{}, err
	}
	if count == 0 {
		newUser.Role = models.AdminRole
	} else {
		newUser.Role = models.UserRole
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	newUser.Password = string(hashedPassword)
	_, err = service.collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return models.User{}, err
	}


	return newUser, nil
}


// login a user

func (service *UserService) GetUser(user models.User) (string, error) {

	var existingUser models.User
	err := service.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", errors.New("invalid password")
	}

	// access jwt secret from .env file
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"Id": existingUser.ID,
	"Username":   existingUser.Username,
	"Role": existingUser.Role,
  })
  
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return jwtToken, nil
}


// promote user to admin

func (service *UserService) PromoteUser(username string) (string, error) {
	var user models.User
	err := service.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}

	user.Role = models.AdminRole
	_, err = service.collection.ReplaceOne(context.Background(), bson.M{"username": username}, user)
	if err != nil {
		return "", err
	}

	return "The user have been promoted", nil
}