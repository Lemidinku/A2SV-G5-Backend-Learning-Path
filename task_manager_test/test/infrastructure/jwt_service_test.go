package infrastructure_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"task_manager_test/domain"
	"task_manager_test/infrastructure"
)

type JWTServiceTestSuite struct {
	suite.Suite
	jwtService *infrastructure.JWTService
	secret     string
}

func (suite *JWTServiceTestSuite) SetupTest() {
	// Set up a custom secret key for the tests
	suite.secret = "mySecretKeyForTests"
	suite.jwtService = infrastructure.NewJWTService(suite.secret)
}

func (suite *JWTServiceTestSuite) TestGenerateToken() {
	id := "123"
	username := "testUser"
	role := domain.Role("admin")

	token, err := suite.jwtService.GenerateToken(id, username, role)
	suite.NoError(err, "GenerateToken() should not return an error")
	suite.NotEmpty(token, "Generated token should not be empty")
}

func (suite *JWTServiceTestSuite) TestValidateToken() {
	id := "123"
	username := "testUser"
	role := domain.AdminRole

	// Generate a token
	token, err := suite.jwtService.GenerateToken(id, username, role)
	suite.NoError(err, "GenerateToken() should not return an error")

	// Validate the token
	validRole, err := suite.jwtService.ValidateToken(token)
	suite.NoError(err, "ValidateToken() should not return an error")
	suite.Equal(string(role), validRole, "ValidateToken() should return the correct role")

	// Test with an invalid token
	_, err = suite.jwtService.ValidateToken("invalidToken")
	suite.Error(err, "ValidateToken() should return an error for an invalid token")
}

func TestJWTServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JWTServiceTestSuite))
}
