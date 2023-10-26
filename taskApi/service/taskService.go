package service

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/google/uuid"
	"time"
)

type TaskService struct {
	amqp *AmqpService
}

func NewTaskService(amqpService *AmqpService) *TaskService {
	return &TaskService{
		amqp: amqpService,
	}
}

func (ts *TaskService) CreateTask(name string) (string, error) {
	id := uuid.New().String()
	task := &model.Task{
		ID:         id,
		Name:       name,
		Status:     model.TASK_CREATED,
		CreateTime: time.Now().String(),
		UpdateTime: time.Now().String(),
	}

	err := ts.amqp.PublishToQueue(task)

	return id, err
}
