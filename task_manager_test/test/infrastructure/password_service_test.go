package infrastructure_test

import (
	"testing"
	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/suite"
	"task_manager_test/infrastructure"
)

type PasswordServiceTestSuite struct {
	suite.Suite
	ps *infrastructure.PasswordService
}

func (suite *PasswordServiceTestSuite) SetupTest() {
	suite.ps = infrastructure.NewPasswordService()
}

func (suite *PasswordServiceTestSuite) TestHashPassword() {
	password := "mySecretPassword"

	hashedPassword, err := suite.ps.HashPassword(password)
	if err != nil {
		suite.T().Fatalf("HashPassword() error = %v", err)
	}

	// Check if hashed password is not empty
	suite.NotEmpty(hashedPassword, "Hashed password should not be empty")

	// Verify the hashed password with bcrypt
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	suite.NoError(err, "HashPassword() failed to verify hashed password")
}

func (suite *PasswordServiceTestSuite) TestIsPasswordCorrect() {
	password := "mySecretPassword"

	hashedPassword, err := suite.ps.HashPassword(password)
	if err != nil {
		suite.T().Fatalf("HashPassword() error = %v", err)
	}

	// Test correct password
	err = suite.ps.IsPasswordCorrect(password, string(hashedPassword))
	suite.NoError(err, "IsPasswordCorrect() should not return an error for correct password")

	// Test incorrect password
	incorrectPassword := "wrongPassword"
	err = suite.ps.IsPasswordCorrect(incorrectPassword, string(hashedPassword))
	suite.Error(err, "IsPasswordCorrect() should return an error for incorrect password")
}

func TestPasswordServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceTestSuite))
}
