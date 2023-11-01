package service

import "github.com/abondar24/TaskDisrtibutor/taskApi/model"

type QueueService interface {
	PublishToQueue(task *model.Task) error
}
