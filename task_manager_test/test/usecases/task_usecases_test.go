package usecases_test

import (
	"errors"
	"task_manager_test/domain"
	"task_manager_test/mocks"
	"task_manager_test/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)


type TaskUsecaseTestSuite struct {
	suite.Suite
	mockTaskRepository *mocks.MockTaskRepository
	taskUsecase        domain.TaskUsecase
}

func (suite *TaskUsecaseTestSuite) SetupSuite() {
	suite.mockTaskRepository = new(mocks.MockTaskRepository)
	suite.taskUsecase = usecases.NewTaskUsecase(suite.mockTaskRepository)
}

func (suite *TaskUsecaseTestSuite) TestGetTasks_Success() {
	// Mock the task repository behavior
	tasks := []domain.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1"},
		{ID: "2", Title: "Task 2", Description: "Description 2"},
	}
	suite.mockTaskRepository.On("GetTasks").Return(tasks, nil)

	// Call the usecase function
	retrievedTasks, err := suite.taskUsecase.GetTasks()

	// Assertions
	suite.NoError(err)
	suite.Len(retrievedTasks, 2)
	suite.Equal("Task 1", retrievedTasks[0].Title)
	suite.Equal("Task 2", retrievedTasks[1].Title)
}

func (suite *TaskUsecaseTestSuite) TestGetTask_Success() {
	// Mock the task repository behavior
	task := domain.Task{ID: "1", Title: "Task 1", Description: "Description 1", DueDate: time.Now(),Status: domain.Pending, }
	suite.mockTaskRepository.On("GetTask", "1").Return(task, nil)

	// Call the usecase function
	retrievedTask, err := suite.taskUsecase.GetTask("1")

	// Assertions
	suite.NoError(err)
	suite.Equal("Task 1", retrievedTask.Title)
}


func (suite *TaskUsecaseTestSuite) TestGetTask_NoTaskFound() {
	// Mock the task repository behavior
	suite.mockTaskRepository.On("GetTask", "1").Return(domain.Task{}, errors.New("task not found"))

	// Call the usecase function
	_, err := suite.taskUsecase.GetTask("1")

	// Assertions
	suite.Error(err)
	// suite.Equal("task not found", err.Error())
}

func (suite *TaskUsecaseTestSuite) TestAddTask_Success() {
	// Mock the task repository behavior
	dueDate := time.Now()
	task := domain.Task{Title: "New Task", Description: "New Description" ,DueDate: dueDate, Status: domain.Pending}
	createdTask := domain.Task{ID: "1", Title: "New Task", Description: "New Description", DueDate: dueDate, Status: domain.Pending}
	suite.mockTaskRepository.On("AddTask", task).Return(createdTask, nil)

	// Call the usecase function
	resultTask, err := suite.taskUsecase.AddTask(task)

	// Assertions
	suite.NoError(err)
	suite.Equal("New Task", resultTask.Title)
	suite.Equal("New Description", resultTask.Description)
	suite.Equal(domain.Pending, resultTask.Status)
	suite.Equal(dueDate, resultTask.DueDate)
}

// TestAddTask_NoDueDate tests the AddTask function when the task does not have a due date
func (suite *TaskUsecaseTestSuite) TestAddTask_NoDueDate() {
	// Mock the task repository behavior
	task := domain.Task{Title: "New Task", Description: "New Description", Status: domain.Pending}
	createdTask := domain.Task{ID: "1", Title: "New Task", Description: "New Description", Status: domain.Pending}
	suite.mockTaskRepository.On("AddTask", task).Return(createdTask, nil)

	// Call the usecase function
	_, err := suite.taskUsecase.AddTask(task)

	// Assertions
	suite.Error(err)
	suite.Equal("due date is required", err.Error())

}

// TestAddTask_NoStatus tests the AddTask function when the task does not have a status
func (suite *TaskUsecaseTestSuite) TestAddTask_NoStatus() {
	// Mock the task repository behavior
	task := domain.Task{Title: "New Task", Description: "New Description", DueDate: time.Now()}
	createdTask := domain.Task{ID: "1", Title: "New Task", Description: "New Description"}
	suite.mockTaskRepository.On("AddTask", task).Return(createdTask, nil)
	
	// Call the usecase function
	_, err := suite.taskUsecase.AddTask(task)

	// Assertions
	suite.Error(err)
	suite.Equal("status is required", err.Error())
}

// TestAddTask_InvalidStatus tests the AddTask function when the task has an invalid status
func (suite *TaskUsecaseTestSuite) TestAddTask_InvalidStatus() {
	// Mock the task repository behavior
	task := domain.Task{Title: "New Task", Description: "New Description", DueDate: time.Now(), Status: "invalid"}
	createdTask := domain.Task{ID: "1", Title: "New Task", Description: "New Description"}
	suite.mockTaskRepository.On("AddTask", task).Return(createdTask, nil)
	
	// Call the usecase function
	_, err := suite.taskUsecase.AddTask(task)

	// Assertions
	suite.Error(err)
	suite.Equal("status must be either completed or pending", err.Error())
}

func (suite *TaskUsecaseTestSuite) TestUpdateTask_Success() {
	// Mock the task repository behavior
	task := domain.Task{Title: "Updated Task", Description: "Updated Description"}
	updatedTask := domain.Task{ID: "1", Title: "Updated Task", Description: "Updated Description"}
	suite.mockTaskRepository.On("UpdateTask", "1", task).Return(updatedTask, nil)

	// Call the usecase function
	resultTask, err := suite.taskUsecase.UpdateTask("1", task)

	// Assertions
	suite.NoError(err)
	suite.Equal("Updated Task", resultTask.Title)
}

func (suite *TaskUsecaseTestSuite) TestUpdateTask_NotFound() {
	// Mock the task repository behavior
	task := domain.Task{Title: "Updated Task", Description: "Updated Description"}
	suite.mockTaskRepository.On("UpdateTask", "2", task).Return(domain.Task{}, errors.New("task not found"))

	// Call the usecase function
	_, err := suite.taskUsecase.UpdateTask("2", task)

	// Assertions
	suite.Error(err)
	suite.Equal("task not found", err.Error())

}




func (suite *TaskUsecaseTestSuite) TestRemoveTask_Success() {
	// Mock the task repository behavior
	suite.mockTaskRepository.On("RemoveTask", "1").Return(nil)

	// Call the usecase function
	err := suite.taskUsecase.RemoveTask("1")

	// Assertions
	suite.NoError(err)
}

func (suite *TaskUsecaseTestSuite) TestRemoveTask_NotFound() {
	// Mock the task repository behavior
	suite.mockTaskRepository.On("RemoveTask", "2").Return(errors.New("task not found"))

	// Call the usecase function
	err := suite.taskUsecase.RemoveTask("2")

	// Assertions
	suite.Error(err)
	suite.Equal("task not found", err.Error())
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}