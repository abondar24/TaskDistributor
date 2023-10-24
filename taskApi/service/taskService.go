package service

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/google/uuid"
)

type TaskService struct {
	//todo - add rmq client
}

func (ts *TaskService) SendTask(name string, status *model.TaskStatus) (string, error) {
	id := uuid.New().String()
	//task :=&model.Task{
	//	ID:         id,
	//	Name:       name,
	//	Status:     status,
	//	CreateTime: time.Now().String(),
	//	UpdateTime: time.Now().String(),
	//}

	//todo - send to rabbit mq

	return id, nil
}
