package repositories

import (
	"context"
	"errors"
	"task_manager_test/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(taskCollection *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		collection: taskCollection,
	}
}

func (repository *TaskRepository) GetTasks() ([]domain.Task, error) {
	cursor, err := repository.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tasks []domain.Task
	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repository *TaskRepository) GetTask(id string) (domain.Task, error) {
	var task domain.Task
	err := repository.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return domain.Task{}, errors.New("task not found")
	} else if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (repository *TaskRepository) AddTask(newTask domain.Task) (domain.Task, error) {
	newTask.ID = uuid.NewString()
	result, err := repository.collection.InsertOne(context.Background(), newTask)
	if err != nil {
		return domain.Task{}, err
	}

	createdTask := domain.Task{}
	err = repository.collection.FindOne(context.Background(), bson.M{"_id": result.InsertedID}).Decode(&createdTask)
	if err != nil {
		return domain.Task{}, err
	}

	return createdTask, nil
}

// UpdateTask updates an existing task in MongoDB
func (repository *TaskRepository) UpdateTask(id string, updatedTask domain.Task) (domain.Task, error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"dueDate":     updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var returnedTask domain.Task
	err := repository.collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&returnedTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("task not found")
		}
		return domain.Task{}, err
	}

	return returnedTask, nil
}

// RemoveTask deletes a task by its ID from MongoDB
func (repository *TaskRepository) RemoveTask(id string) error {
	filter := bson.M{"_id": id}

	result, err := repository.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
