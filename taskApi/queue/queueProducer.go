package queue

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
)

type Producer interface {
	PublishToQueue(task *data.Task) error
}
