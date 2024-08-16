package mocks

import (
	"task_manager_test/domain"

	"github.com/stretchr/testify/mock"
)

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