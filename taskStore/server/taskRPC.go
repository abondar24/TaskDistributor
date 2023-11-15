package server

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskData/rpc"
	"github.com/abondar24/TaskDistributor/taskStore/service"
	"log"
	"net/http"
)

type TaskRPC struct {
	taskService service.TaskService
}

func NewTaskRPC(taskService service.TaskService) *TaskRPC {

	return &TaskRPC{
		taskService: taskService,
	}
}

// GetTask
// @Summary Get tasks
// @Description fetch task by id,status or history. By id or status data.task is returned
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param rpcRequest body rpc.TaskRPCRequest  true "RPC Request"
// @Param rpcRequest body rpc.TaskRPCRequest  true "RPC Request for status"
// @Success 200 {object} data.Task
// @Success 200 {object} data.TaskHistory
// @BadRequest 400
// @Router /rpc [post]
func (tr *TaskRPC) GetTask(r *http.Request, id *string, response *data.Task) error {

	log.Printf("Fetching task by id %s\n", *id)

	res, err := tr.taskService.GetTaskById(id)
	if err != nil {
		return err
	}

	*response = *res

	return nil
}

func (tr *TaskRPC) GetTasksByStatus(r *http.Request, args *rpc.StatusArgs, response *[]*data.Task) error {
	log.Printf("Fetching tasks by status %s\n", *args.Status)

	res, err := tr.taskService.GetTasksByStatus(args.Status, args.Offset, args.Limit)
	if err != nil {
		return err
	}

	*response = res

	return nil
}

func (tr *TaskRPC) GetTaskHistory(r *http.Request, id *string, response *data.TaskHistory) error {
	log.Printf("Fetching task history for id %s\n", *id)

	res, err := tr.taskService.GetTaskHistory(id)
	if err != nil {
		return err
	}

	*response = *res

	return nil
}
