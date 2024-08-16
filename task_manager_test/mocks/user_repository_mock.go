package mocks

import (
	"task_manager_test/domain"

	"github.com/stretchr/testify/mock"
)

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