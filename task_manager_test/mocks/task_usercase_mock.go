package mocks

import (
	"task_manager_test/domain"

	"github.com/stretchr/testify/mock"
)

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