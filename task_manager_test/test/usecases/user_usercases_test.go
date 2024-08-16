package usecases_test

import (
	"errors"
	"task_manager_test/domain"
	"task_manager_test/mocks"
	"task_manager_test/usecases"
	"testing"

	"github.com/stretchr/testify/suite"
)



type UserUsecaseTestSuite struct {
	suite.Suite
	userRepositoryMock  *mocks.UserRepositoryMock
	passwordService *mocks.PasswordServiceMock
	jwtService *mocks.JWTServiceMock
	userUsecase         domain.UserUsecase
}



func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.userRepositoryMock = new(mocks.UserRepositoryMock)
	suite.passwordService = new(mocks.PasswordServiceMock)
	suite.jwtService = new(mocks.JWTServiceMock)
	suite.userUsecase = usecases.NewUserUsecase(suite.userRepositoryMock, suite.passwordService, suite.jwtService)
}

func (suite *UserUsecaseTestSuite) TestRegisterUser_Success() {
	mockUser := domain.User{Username: "test_user", Password: "hashed_password"}
	// Mock repository behavior
	suite.userRepositoryMock.On("RegisterUser", mockUser).Return(mockUser, nil)
	// Mock password service behavior
	suite.passwordService.On("HashPassword", mockUser.Password).Return([]byte("hashed_password"), nil)

	// Call the use case method
	createdUser, err := suite.userUsecase.RegisterUser(mockUser)

	// Assertions
	suite.NoError(err)
	suite.Equal(mockUser.Username, createdUser.Username)
	suite.Equal("hashed_password", createdUser.Password)
	suite.userRepositoryMock.AssertExpectations(suite.T())
}

// no password
func (suite *UserUsecaseTestSuite) TestRegisterUser_NoPassword() {
	
	mockUser := domain.User{Username: "test_user", Password: ""}
	// Mock repository behavior
	suite.userRepositoryMock.On("RegisterUser", mockUser).Return(mockUser, nil)

	// Call the use case method
	_, err := suite.userUsecase.RegisterUser(mockUser)

	// Assertions
	suite.Error(err)
	suite.Equal("password is required", err.Error())
}

func (suite *UserUsecaseTestSuite) TestRegisterUser_ShortPassword() {
	
	mockUser := domain.User{Username: "test_user", Password: "lemi"}
	// Mock repository behavior
	suite.userRepositoryMock.On("RegisterUser", mockUser).Return(mockUser, nil)

	// Call the use case method
	_, err := suite.userUsecase.RegisterUser(mockUser)

	// Assertions
	suite.Error(err)
	suite.Equal("password must be at least 6 characters", err.Error())
}

func (suite *UserUsecaseTestSuite) TestRegisterUser_NoUsername() {
	
	mockUser := domain.User{Username: "", Password: "asdfasdf"}
	// Mock repository behavior
	suite.userRepositoryMock.On("RegisterUser", mockUser).Return(mockUser, nil)

	// Call the use case method
	_, err := suite.userUsecase.RegisterUser(mockUser)

	// Assertions
	suite.Error(err)
	suite.Equal("username is required", err.Error())
}

func (suite *UserUsecaseTestSuite) TestRegisterUser_ShortUsername() {
	
	mockUser := domain.User{Username: "no", Password: "lemidinku"}
	// Mock repository behavior
	suite.userRepositoryMock.On("RegisterUser", mockUser).Return(mockUser, nil)

	// Call the use case method
	_, err := suite.userUsecase.RegisterUser(mockUser)

	// Assertions
	suite.Error(err)
	suite.Equal("username must be at least 3 characters", err.Error())
}





func (suite *UserUsecaseTestSuite) TestGetUser_Success() {
	mockUser := domain.User{Username: "test_user", Password: "password123"}
	mockExistingUser := domain.User{Username: "test_user", Password: "password123"}
	suite.userRepositoryMock.On("GetUser", mockUser).Return(mockExistingUser, nil)
	suite.passwordService.On("IsPasswordCorrect", mockUser.Password, mockExistingUser.Password).Return(nil)
	suite.jwtService.On("GenerateToken", mockExistingUser.ID, mockExistingUser.Username, mockExistingUser.Role).Return("jwt_token", nil)
	// call the use case method
	jwtToken, err := suite.userUsecase.GetUser(mockUser)

	// Assertions
	suite.NoError(err)
	suite.NotNil(jwtToken)
}

func (suite *UserUsecaseTestSuite) TestGetUser_WrongPassword() {
	mockUser := domain.User{Username: "test_user", Password: "password123"}
	mockExistingUser := domain.User{Username: "test_user", Password: "password123"}

	suite.userRepositoryMock.On("GetUser", mockUser).Return(mockExistingUser, nil)
	suite.passwordService.On("IsPasswordCorrect", mockUser.Password, mockExistingUser.Password).Return(errors.New("invalid password"))

	// Call the use case method
	_, err := suite.userUsecase.GetUser(mockUser)

	// Assertions
	suite.Error(err)
	suite.Equal("invalid password", err.Error())

}

func (suite *UserUsecaseTestSuite) TestGetUser_NotFound() {
	mockUser := domain.User{Username: "no_user", Password: "password123"}
	suite.userRepositoryMock.On("GetUser", mockUser).Return(domain.User{}, errors.New("user not found"))
	// call the use case method
	_, err := suite.userUsecase.GetUser(mockUser)

	// Assertions
	suite.Error(err)
	suite.Equal("user not found", err.Error())

}

func (suite *UserUsecaseTestSuite) TestPromoteUser_Success() {
	mockUsername := "test_user"
	suite.userRepositoryMock.On("PromoteUser", mockUsername).Return("The user have been promoted", nil)

	// Call the use case method
	message, err := suite.userUsecase.PromoteUser(mockUsername)

	// Assertions
	suite.NoError(err)
	suite.Equal("The user have been promoted", message)
	
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
