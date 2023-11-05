package service

import (
	"github.com/abondar24/TaskDistributor/taskApi/queue"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/google/uuid"
	"time"
)

type TaskCommandService struct {
	amqp queue.Producer
}

func NewTaskService(amqpService queue.Producer) *TaskCommandService {
	return &TaskCommandService{
		amqp: amqpService,
	}
}

func (ts *TaskCommandService) CreateTask(name *string) (string, error) {
	id := uuid.New().String()
	task := &data.Task{
		ID:         id,
		Name:       *name,
		Status:     data.TASK_CREATED,
		CreateTime: time.Now().String(),
		UpdateTime: time.Now().String(),
	}

	err := ts.amqp.PublishToQueue(task)

	return id, err
}

func (ts *TaskCommandService) UpdateTask(id *string, completed *bool) error {
	task := data.Task{
		ID:         *id,
		Status:     data.TASK_UPDATED,
		UpdateTime: time.Now().String(),
	}

	if *completed {
		task.Status = data.TASK_COMPLETED
	}

	err := ts.amqp.PublishToQueue(&task)

	return err
}

func (ts *TaskCommandService) DeleteTask(id *string) error {
	task := &data.Task{
		ID:         *id,
		Status:     data.TASK_DELETED,
		UpdateTime: time.Now().String(),
	}

	err := ts.amqp.PublishToQueue(task)

	return err
}
