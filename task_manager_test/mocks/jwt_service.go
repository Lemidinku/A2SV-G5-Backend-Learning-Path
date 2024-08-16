package mocks

import (
	"task_manager_test/domain"

	"github.com/stretchr/testify/mock"
)

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