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


