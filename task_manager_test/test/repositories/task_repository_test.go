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

type TaskRepositoryTestSuite struct {
	suite.Suite
	client     *mongo.Client
	collection *mongo.Collection
	repo       *repositories.TaskRepository
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	suite.NoError(err)

	// Wait for the connection
	err = client.Ping(context.Background(), nil)
	suite.NoError(err)

	suite.client = client
	suite.collection = client.Database("task_manager_test").Collection("tasks")
	suite.repo = repositories.NewTaskRepository(suite.collection)
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	// Disconnect from MongoDB
	err := suite.client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *TaskRepositoryTestSuite) SetupTest() {
	// Clean the collection before each test
	err := suite.collection.Drop(context.Background())
	suite.NoError(err)
}

func (suite *TaskRepositoryTestSuite) TestGetTasks_Success() {
	// Insert test data
	_, err := suite.collection.InsertMany(context.Background(), []interface{}{
		bson.M{"_id": "1", "title": "Task 1", "description": "Description 1"},
		bson.M{"_id": "2", "title": "Task 2", "description": "Description 2"},
	})
	suite.NoError(err)

	// Retrieve tasks
	tasks, err := suite.repo.GetTasks()

	// Assertions
	suite.NoError(err)
	suite.Len(tasks, 2)
	suite.Equal("Task 1", tasks[0].Title)
	suite.Equal("Task 2", tasks[1].Title)
}

func (suite *TaskRepositoryTestSuite) TestGetTasks_NoTasks() {
	// Retrieve tasks
	tasks, err := suite.repo.GetTasks()
	// Assertions
	suite.NoError(err)
	suite.Len(tasks, 0)

}

func (suite *TaskRepositoryTestSuite) TestGetTask() {
	// Insert a test task
	_, err := suite.collection.InsertOne(context.Background(), bson.M{
		"_id":         "1",
		"title":       "Task 1",
		"description": "Description 1",
	})
	suite.NoError(err)

	// Retrieve task by ID
	task, err := suite.repo.GetTask("1")

	// Assertions
	suite.NoError(err)
	suite.Equal("Task 1", task.Title)
	suite.Equal("Description 1", task.Description)
}

// task not found test
func (suite *TaskRepositoryTestSuite) TestGetTaskNotFound() {
	// Retrieve task by ID
	_, err := suite.repo.GetTask("1")
	suite.Error(err)
	// check the error message
	suite.Equal("task not found", err.Error())
}


func (suite *TaskRepositoryTestSuite) TestAddTask() {
	// Add a new task
	newTask := domain.Task{Title: "New Task", Description: "New Description"}
	createdTask, err := suite.repo.AddTask(newTask)

	// Assertions
	suite.NoError(err)
	suite.NotEmpty(createdTask.ID)
	suite.Equal("New Task", createdTask.Title)

	// Check the task was added
	var task domain.Task
	err = suite.collection.FindOne(context.Background(), bson.M{"_id": createdTask.ID}).Decode(&task)
	suite.NoError(err)
	suite.NotNil(task.ID)
	suite.Equal("New Task", task.Title)
	suite.Equal("New Description", task.Description)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask() {
	// Insert a test task
	_, err := suite.collection.InsertOne(context.Background(), bson.M{
		"_id":         "1",
		"title":       "Task 1",
		"description": "Description 1",
	})
	suite.NoError(err)

	// Update the task
	updatedTask := domain.Task{Title: "Updated Task", Description: "Updated Description"}
	task, err := suite.repo.UpdateTask("1", updatedTask)

	// Assertions
	suite.NoError(err)
	suite.Equal("Updated Task", task.Title)
	suite.Equal("Updated Description", task.Description)
}

func (suite *TaskRepositoryTestSuite) TestUpdateTask_NotFound() {
	// Update the task
	updatedTask := domain.Task{Title: "Updated Task", Description: "Updated Description"}
	_, err := suite.repo.UpdateTask("1", updatedTask)

	// Assertions
	suite.Error(err)
	suite.Equal("task not found", err.Error())
}

func (suite *TaskRepositoryTestSuite) TestRemoveTask() {
	// Insert a test task
	_, err := suite.collection.InsertOne(context.Background(), bson.M{
		"_id": "1", "title": "Task 1", "description": "Description 1",
	})
	suite.NoError(err)

	// Remove the task
	err = suite.repo.RemoveTask("1")

	// Assertions
	suite.NoError(err)

	// Verify task was removed
	count, err := suite.collection.CountDocuments(context.Background(), bson.M{"_id": "1"})
	suite.NoError(err)
	suite.Equal(int64(0), count)
}

func (suite *TaskRepositoryTestSuite) TestRemoveTask_NotFound() {
	// Remove a task that does not exist
	err := suite.repo.RemoveTask("1")

	// Assertions
	suite.Error(err)
	suite.Equal("task not found", err.Error())
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
