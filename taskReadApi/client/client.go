package client

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskData/rpc"
)

type Client interface {
	GetTask(id *string) (*data.Task, error)

	GetTasksByStatus(args *rpc.StatusArgs) (*[]*data.Task, error)

	GetTaskHistory(id *string) (*data.TaskHistory, error)
}
