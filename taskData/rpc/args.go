package rpc

import "github.com/abondar24/TaskDistributor/taskData/data"

type StatusArgs struct {
	Status *data.TaskStatus `json:"status"`
	Offset *int             `json:"offset"`
	Limit  *int             `json:"limit"`
}
