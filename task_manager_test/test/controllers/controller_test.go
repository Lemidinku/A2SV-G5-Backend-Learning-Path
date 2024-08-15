// package controllers

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"task_manager_test/delivery/controllers"
// 	"task_manager_test/domain"
// 	"task_manager_test/mocks"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type ControllerTestSuite struct {
// 	suite.Suite
// 	taskController *controllers.TaskController
// 	taskUsecaseMock *mocks.TaskUsecaseMock
// 	userUsecaseMock *mocks.UserUsecaseMock
// 	router          *gin.Engine
// }

// func (suite *ControllerTestSuite) SetupTest() {
// 	suite.taskUsecaseMock = new(mocks.TaskUsecaseMock)
// 	suite.userUsecaseMock = new(mocks.UserUsecaseMock)
// 	suite.taskController = controllers.NewTaskController(suite.taskUsecaseMock, suite.userUsecaseMock)

// 	suite.router = gin.Default()
// }

// func (suite *ControllerTestSuite) TestGetTasks() {
// 	// Mock the usecase response
// 	suite.taskUsecaseMock.On("GetTasks").Return([]domain.Task{{ID: "1", Title: "Test Task"}}, nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("GET", "/tasks", nil)
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.GET("/tasks", suite.taskController.GetTasks)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "Test Task")
// }

// func (suite *ControllerTestSuite) TestGetTask() {
// 	// Mock the usecase response
// 	suite.taskUsecaseMock.On("GetTask", "1").Return(domain.Task{ID: "1", Title: "Test Task"}, nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("GET", "/tasks/1", nil)
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.GET("/tasks/:id", suite.taskController.GetTask)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "Test Task")
// }

// func (suite *ControllerTestSuite) TestAddTask() {
// 	taskJSON := `{"title":"New Task","description":"Task description"}`

// 	// Mock the usecase response
// 	suite.taskUsecaseMock.On("AddTask", mock.Anything).Return(domain.Task{ID: "1", Title: "New Task"}, nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("POST", "/tasks", strings.NewReader(taskJSON))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.POST("/tasks", suite.taskController.AddTask)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusCreated, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "New Task")
// }

// func (suite *ControllerTestSuite) TestRemoveTask() {
// 	// Mock the usecase response
// 	suite.taskUsecaseMock.On("RemoveTask", "1").Return(nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.DELETE("/tasks/:id", suite.taskController.RemoveTask)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "Task removed")
// }

// func (suite *ControllerTestSuite) TestUpdateTask() {
// 	taskJSON := `{"title":"Updated Task","description":"Updated description"}`

// 	// Mock the usecase response
// 	suite.taskUsecaseMock.On("UpdateTask", "1", mock.Anything).Return(domain.Task{ID: "1", Title: "Updated Task"}, nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("PUT", "/tasks/1", strings.NewReader(taskJSON))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.PUT("/tasks/:id", suite.taskController.UpdateTask)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "Updated Task")
// }

// func (suite *ControllerTestSuite) TestRegister() {
// 	userJSON := `{"username":"newuser","password":"password123"}`

// 	// Mock the usecase response
// 	suite.userUsecaseMock.On("RegisterUser", mock.Anything).Return(domain.User{Username: "newuser"}, nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("POST", "/register", strings.NewReader(userJSON))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.POST("/register", suite.taskController.Register)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "newuser")
// }

// func (suite *ControllerTestSuite) TestLogin() {
// 	userJSON := `{"username":"existinguser","password":"password123"}`

// 	// Mock the usecase response
// 	suite.userUsecaseMock.On("GetUser", mock.Anything).Return("mockToken", nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("POST", "/login", strings.NewReader(userJSON))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.POST("/login", suite.taskController.Login)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "User logged in successfully")
// 	assert.Contains(suite.T(), w.Body.String(), "mockToken")
// }

// func (suite *ControllerTestSuite) TestPromoteUser() {
// 	// Mock the usecase response
// 	suite.userUsecaseMock.On("PromoteUser", "existinguser").Return("User promoted to admin", nil)

// 	// Create a request to pass to the handler
// 	req, _ := http.NewRequest("PUT", "/promote/existinguser", nil)
// 	w := httptest.NewRecorder()

// 	// Register the handler and make the request
// 	suite.router.PUT("/promote/:username", suite.taskController.PromoteUser)
// 	suite.router.ServeHTTP(w, req)

// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	assert.Contains(suite.T(), w.Body.String(), "User promoted to admin")
// }

// func TestControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(ControllerTestSuite))
// }

package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task_manager_test/delivery/controllers"
	"task_manager_test/domain"
	"task_manager_test/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	taskController  *controllers.TaskController
	taskUsecaseMock *mocks.TaskUsecaseMock
	userUsecaseMock *mocks.UserUsecaseMock
	testingServer   *httptest.Server
}

func (suite *ControllerTestSuite) SetupSuite() {
	// Initialize mock usecases
	suite.taskUsecaseMock = new(mocks.TaskUsecaseMock)
	suite.userUsecaseMock = new(mocks.UserUsecaseMock)
	suite.taskController = controllers.NewTaskController(suite.taskUsecaseMock, suite.userUsecaseMock)

	// Setup router and routes
	router := gin.Default()

	// Task routes
	router.GET("/tasks", suite.taskController.GetTasks)
	router.GET("/tasks/:id", suite.taskController.GetTask)
	router.POST("/tasks", suite.taskController.AddTask)
	router.PUT("/tasks/:id", suite.taskController.UpdateTask)
	router.DELETE("/tasks/:id", suite.taskController.RemoveTask)

	// User routes
	router.POST("/register", suite.taskController.Register)
	router.POST("/login", suite.taskController.Login)
	router.PATCH("/promote/:username", suite.taskController.PromoteUser)

	// Create a test server
	suite.testingServer = httptest.NewServer(router)
}

func (suite *ControllerTestSuite) TearDownSuite() {
	// Close the test server after the suite has finished running
	suite.testingServer.Close()
}

func (suite *ControllerTestSuite) TestGetTasks() {
	suite.taskUsecaseMock.On("GetTasks").Return([]domain.Task{{ID: "1", Title: "Test Task"}}, nil)

	resp, err := http.Get(suite.testingServer.URL + "/tasks")
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestGetTask() {
	suite.taskUsecaseMock.On("GetTask", "1").Return(domain.Task{ID: "1", Title: "Test Task"}, nil)

	resp, err := http.Get(suite.testingServer.URL + "/tasks/1")
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestCreateTask() {
	taskJSON := `{"title":"New Task","description":"Task description"}`
	suite.taskUsecaseMock.On("AddTask", mock.Anything).Return(domain.Task{ID: "1", Title: "New Task"}, nil)

	resp, err := http.Post(suite.testingServer.URL+"/tasks", "application/json", strings.NewReader(taskJSON))
	suite.NoError(err)
	suite.Equal(http.StatusCreated, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestUpdateTask() {
	taskJSON := `{"title":"Updated Task","description":"Updated description"}`
	suite.taskUsecaseMock.On("UpdateTask", "1", mock.Anything).Return(domain.Task{ID: "1", Title: "Updated Task"}, nil)

	req, _ := http.NewRequest("PUT", suite.testingServer.URL+"/tasks/1", strings.NewReader(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestDeleteTask() {
	suite.taskUsecaseMock.On("RemoveTask", "1").Return(nil)

	req, _ := http.NewRequest("DELETE", suite.testingServer.URL+"/tasks/1", nil)
	resp, err := http.DefaultClient.Do(req)

	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestSignup() {
	userJSON := `{"username":"newuser","password":"password123"}`
	suite.userUsecaseMock.On("RegisterUser", mock.Anything).Return(domain.User{Username: "newuser"}, nil)

	resp, err := http.Post(suite.testingServer.URL+"/register", "application/json", strings.NewReader(userJSON))
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestLogin() {
	userJSON := `{"username":"existinguser","password":"password123"}`
	suite.userUsecaseMock.On("GetUser", mock.Anything).Return("mockToken", nil)

	resp, err := http.Post(suite.testingServer.URL+"/login", "application/json", strings.NewReader(userJSON))
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *ControllerTestSuite) TestPromoteUser() {
	suite.userUsecaseMock.On("PromoteUser", "existinguser").Return("User promoted to admin", nil)

	req, _ := http.NewRequest("PATCH", suite.testingServer.URL+"/promote/existinguser", nil)
	resp, err := http.DefaultClient.Do(req)

	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
