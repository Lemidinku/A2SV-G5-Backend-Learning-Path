package services

import (
	"context"
	"errors"
	"task_manager2/database"
	"task_manager2/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService() *TaskService {
	return &TaskService{
		collection: database.TaskCollection,
	}
}

func (service *TaskService) GetTasks() ([]models.Task, error) {
	cursor, err := service.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var tasks []models.Task
	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (service *TaskService) GetTask(id string) (models.Task, error) {

	var task models.Task
	err := service.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	} else if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (service *TaskService) AddTask(newTask models.Task) (models.Task, error) {
	newTask.ID = uuid.NewString()
	result, err := service.collection.InsertOne(context.Background(), newTask)
	if err != nil {
		return models.Task{}, err
	}

	createdTask := models.Task{}
	err = service.collection.FindOne(context.Background(), bson.M{"_id": result.InsertedID}).Decode(&createdTask)
	if err != nil {
		return models.Task{}, err
	}

	return createdTask, nil
}

// UpdateTask updates an existing task in MongoDB
func (service *TaskService) UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"dueDate":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var returnedTask models.Task
	err := service.collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&returnedTask)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	return updatedTask, nil
}

// RemoveTask deletes a task by its ID from MongoDB
func (service *TaskService) RemoveTask(id string) error {
	filter := bson.M{"_id": id}

	result, err := service.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
