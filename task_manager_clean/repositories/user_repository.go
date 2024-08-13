package repositories

import (
	"context"
	"errors"
	"task_manager_clean/domain"
	"task_manager_clean/infrastructure"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(userCollection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: userCollection,
	}
}

// register a new user
func (repo *UserRepository) RegisterUser(newUser domain.User) (domain.User, error) {
	// check if username already exists
	userNameCount, err := repo.collection.CountDocuments(context.Background(), bson.M{"username": newUser.Username})
	if err != nil {
		return domain.User{}, err
	}
	if userNameCount > 0 {
		return domain.User{}, errors.New("username already taken")
	}

	newUser.ID = uuid.NewString()

	// check if user is the first user to be registered
	count, err := repo.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return domain.User{}, err
	}
	if count == 0 {
		newUser.Role = domain.AdminRole
	} else {
		newUser.Role = domain.UserRole
	}

	hashedPassword, err := infrastructure.HashPassword(newUser.Password)
	if err != nil {
		return domain.User{}, err
	}
	newUser.Password = string(hashedPassword)

	_, err = repo.collection.InsertOne(context.Background(), newUser)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

// login a user
func (repo *UserRepository) GetUser(user domain.User) (string, error) {
	var existingUser domain.User
	err := repo.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}
	err = infrastructure.IsPasswordCorrect(user.Password, existingUser.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}



	jwtToken, err := infrastructure.GenerateToken(existingUser.ID, existingUser.Username, existingUser.Role)

	if err != nil {
		return "", errors.New("internal server error")
	}

	return jwtToken, nil
}

// promote user to admin
func (repo *UserRepository) PromoteUser(username string) (string, error) {
	var user domain.User
	err := repo.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return "", errors.New("user not found")
	} else if err != nil {
		return "", err
	}

	user.Role = domain.AdminRole
	_, err = repo.collection.ReplaceOne(context.Background(), bson.M{"username": username}, user)
	if err != nil {
		return "", err
	}

	return "The user has been promoted", nil
}
