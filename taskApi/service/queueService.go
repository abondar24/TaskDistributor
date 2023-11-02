package service

import "github.com/abondar24/TaskDistributor/taskApi/model"

type QueueService interface {
	PublishToQueue(task *model.Task) error
}
