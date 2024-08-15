package usecases

import (
	"errors"
	"task_manager_test/domain"
)

// TaskUsecase represent the task's usecases
type UserUsecase struct {
	userRepository domain.UserRepository
	passwordService domain.PasswordService
	jwtService domain.JWTService
}

// NewTaskUsecase will create new an taskUsecase object representation of domain.TaskUsecase interface
func NewUserUsecase(ur domain.UserRepository, ps domain.PasswordService, jwtServ domain.JWTService) domain.UserUsecase {
	return &UserUsecase{
		userRepository: ur,
		passwordService: ps,
		jwtService: jwtServ,
	}
}

// GetTasks will get all tasks
func (tu *UserUsecase) RegisterUser(user domain.User) (domain.User, error) {
	// validate the user
	err := user.ValidateUser()
	if err != nil {
		return domain.User{}, err
	}
	// instantiate the password service
	hashedPassword, err := tu.passwordService.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = string(hashedPassword)
	return tu.userRepository.RegisterUser(user)
}	

// GetTask will get a task by its ID
func (tu *UserUsecase) GetUser(user domain.User) (string, error) {
	existingUser, err := tu.userRepository.GetUser(user)
	if err != nil {
		return "", err
	}
	// instantiate the password service
	err = tu.passwordService.IsPasswordCorrect(user.Password, existingUser.Password)
	if err != nil {
		return "", errors.New("invalid password")
	}
	jwtToken, err := tu.jwtService.GenerateToken(existingUser.ID, existingUser.Username, existingUser.Role)

	if err != nil {
		return "", errors.New("internal server error")
	}

	return jwtToken, nil
}

// AddTask will add new task
func (tu *UserUsecase) PromoteUser(username string) (string, error) {
	return tu.userRepository.PromoteUser(username)

}


