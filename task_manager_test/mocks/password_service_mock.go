package mocks

import (
	"github.com/stretchr/testify/mock"
)

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