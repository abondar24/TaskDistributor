package queue

import "github.com/abondar24/TaskDistributor/taskApi/model"

type Producer interface {
	PublishToQueue(task *model.Task) error
}
