package usecases

import (
	"task_manager_clean/domain"
)

// TaskUsecase represent the task's usecases
type UserUsecase struct {
	userRepository domain.UserRepository
}

// NewTaskUsecase will create new an taskUsecase object representation of domain.TaskUsecase interface
func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository: ur,
	}
}

// GetTasks will get all tasks
func (tu *UserUsecase) RegisterUser(user domain.User) (domain.User, error) {
	return tu.userRepository.RegisterUser(user)
}	

// GetTask will get a task by its ID
func (tu *UserUsecase) GetUser(user domain.User) (string, error) {
	return tu.userRepository.GetUser(user)
}

// AddTask will add new task
func (tu *UserUsecase) PromoteUser(username string) (string, error) {
	return tu.userRepository.PromoteUser(username)

}


