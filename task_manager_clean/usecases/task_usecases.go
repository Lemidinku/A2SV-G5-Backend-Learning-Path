package usecases

import (
	"task_manager_clean/domain"
)

// TaskUsecase represent the task's usecases
type TaskUsecase struct {
	taskRepository domain.TaskRepository
}


// NewTaskUsecase will create new an taskUsecase object representation of domain.TaskUsecase interface
func NewTaskUsecase(tr domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{
		taskRepository: tr,
	}
}

// GetTasks will get all tasks
func (tu *TaskUsecase) GetTasks() ([]domain.Task, error) {
	return tu.taskRepository.GetTasks()
}

// GetTask will get a task by its ID
func (tu *TaskUsecase) GetTask(id string) (domain.Task, error) {
	return tu.taskRepository.GetTask(id)
}


// AddTask will add new task
func (tu *TaskUsecase) AddTask(task domain.Task) (domain.Task, error) {
	err := task.ValidateTask()
	if err != nil {
		return domain.Task{}, err
	}

	return tu.taskRepository.AddTask(task)
}

// UpdateTask will update task's data
func (tu *TaskUsecase) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	return tu.taskRepository.UpdateTask(id, task)
}

// RemoveTask will remove task by its ID
func (tu *TaskUsecase) RemoveTask(id string) error {
	return tu.taskRepository.RemoveTask(id)
}

