package mocks

import (
	"task_manager_test/domain"

	"github.com/stretchr/testify/mock"
)

// MockTaskRepository is a mock implementation of the TaskRepository interface
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) AddTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskRepository) RemoveTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}


// MockUserRepository is a mock implementation of the UserRepository interface
type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) RegisterUser(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUser(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepositoryMock) PromoteUser(username string) (string, error) {
	args := m.Called(username)
	return args.String(0), args.Error(1)
}


// MockPasswordService is a mock implementation of the PasswordService interface
type PasswordServiceMock struct {
	mock.Mock
}

func (m *PasswordServiceMock) HashPassword(password string) ([]byte, error) {
	args := m.Called(password)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *PasswordServiceMock) IsPasswordCorrect(password, hashedPassword string) error {
	args := m.Called(password, hashedPassword)
	return args.Error(0)
}


// MockJWTService is a mock implementation of the JWTService interface
type JWTServiceMock struct {
	mock.Mock
}

func (m *JWTServiceMock) GenerateToken(id, username string, role domain.Role) (string, error) {
	args := m.Called(id, username, role)
	return args.String(0), args.Error(1)
}

func (m *JWTServiceMock) ValidateToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}



// MockUserUsecase is a mock implementation of the UserUsecase interface
type UserUsecaseMock struct {
	mock.Mock
}

func (m *UserUsecaseMock) RegisterUser(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserUsecaseMock) GetUser(user domain.User) (string, error) {
	args := m.Called(user)
	return args.Get(0).(string), args.Error(1)
}

func (m *UserUsecaseMock) PromoteUser(username string) (string, error) {
	args := m.Called(username)
	return args.Get(0).(string), args.Error(1)
}



// TaskUsecaseMock is a mock implementation of the TaskUsecase interface
type TaskUsecaseMock struct {
	mock.Mock
}

func (m *TaskUsecaseMock) GetTasks() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *TaskUsecaseMock) GetTask(id string) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *TaskUsecaseMock) AddTask(task domain.Task) (domain.Task, error) {
	args := m.Called(task)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *TaskUsecaseMock) RemoveTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *TaskUsecaseMock) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(domain.Task), args.Error(1)
}