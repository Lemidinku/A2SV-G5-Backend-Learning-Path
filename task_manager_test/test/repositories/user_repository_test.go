package repositories_test

import (
	"context"
	"task_manager_test/domain"
	"task_manager_test/repositories"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
	repo       *repositories.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.NoError(err)

	// Wait for the connection
	err = client.Ping(context.Background(), nil)
	suite.NoError(err)

	suite.client = client
	suite.collection = client.Database("task_manager_test").Collection("users")
	suite.repo = repositories.NewUserRepository(suite.collection)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	// Disconnect from MongoDB
	err := suite.client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	// Clean the collection before each test
	err := suite.collection.Drop(context.Background())
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestRegisterUser() {
	// Register a new user
	newUser := domain.User{Username: "testuser", Password: "password123"}
	createdUser, err := suite.repo.RegisterUser(newUser)

	// Assertions
	suite.NoError(err)
	suite.NotEmpty(createdUser.ID)
	suite.Equal("testuser", createdUser.Username)
	suite.Equal("password123", createdUser.Password) 
	suite.Equal(domain.AdminRole, createdUser.Role) // First user should be an admin

	// Check the user was added
	var user domain.User
	err = suite.collection.FindOne(context.Background(), bson.M{"_id": createdUser.ID}).Decode(&user)
	suite.NoError(err)
	suite.Equal("testuser", user.Username)
}

// username already taken
func (suite *UserRepositoryTestSuite) TestRegisterUser_UsernameTaken() {
	// Insert a user with the same username
	_, err := suite.collection.InsertOne(context.Background(), domain.User{Username: "testuser"})
	suite.NoError(err)

	// Attempt to register a new user with the same username
	newUser := domain.User{Username: "testuser", Password: "password123"}
	_, err = suite.repo.RegisterUser(newUser)

	// Assertions
	suite.Error(err)
	suite.Equal("username already taken", err.Error())
}

func (suite *UserRepositoryTestSuite) TestGetUser_Success() {
	// Insert a user into the collection
	_, err := suite.collection.InsertOne(context.Background(), domain.User{Username: "testuser", Password: "password123"})
	suite.NoError(err)

	// Attempt to retrieve the user
	user, err := suite.repo.GetUser(domain.User{Username: "testuser"})

	// Assertions
	suite.NoError(err)
	suite.Equal("testuser", user.Username)
}

func (suite *UserRepositoryTestSuite) TestGetUser_NotFound() {
	// Attempt to retrieve a user that does not exist
	_, err := suite.repo.GetUser(domain.User{Username: "nonexistentuser"})

	// Assertions
	suite.Error(err)
	suite.Equal("user not found", err.Error())
}

func (suite *UserRepositoryTestSuite) TestPromoteUser_Success() {
	// Insert a user into the collection
	_, err := suite.collection.InsertOne(context.Background(), domain.User{Username: "testuser", Role: domain.UserRole})
	suite.NoError(err)

	// Promote the user to admin
	message, err := suite.repo.PromoteUser("testuser")

	// Assertions
	suite.NoError(err)
	suite.Equal("The user has been promoted", message)

	// Verify the user's role was updated
	var user domain.User
	err = suite.collection.FindOne(context.Background(), bson.M{"username": "testuser"}).Decode(&user)
	suite.NoError(err)
	suite.Equal(domain.AdminRole, user.Role)
}

func (suite *UserRepositoryTestSuite) TestPromoteUser_NotFound() {
	// Attempt to promote a user that does not exist
	_, err := suite.repo.PromoteUser("nonexistentuser")

	// Assertions
	suite.Error(err)
	suite.Equal("user not found", err.Error())
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
