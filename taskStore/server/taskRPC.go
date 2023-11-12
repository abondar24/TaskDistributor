package server

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/service"
)

type TaskRPC struct {
	taskService service.TaskService
}

func NewTaskRPC(taskService service.TaskService) *TaskRPC {

	return &TaskRPC{
		taskService: taskService,
	}
}

func (tr *TaskRPC) GetTask(id *string, result *data.Task) error {
	res, err := tr.taskService.GetTaskById(id)
	if err != nil {
		return err
	}

	*result = *res

	return nil
}

func (tr *TaskRPC) GetTasksByStatus(status *data.TaskStatus, offset *int, limit *int, result *[]*data.Task) error {
	res, err := tr.taskService.GetTasksByStatus(status, offset, limit)
	if err != nil {
		return err
	}

	*result = res

	return nil
}

func (tr *TaskRPC) GetTaskHistory(id *string, result *data.TaskHistory) error {
	res, err := tr.taskService.GetTaskHistory(id)
	if err != nil {
		return err
	}

	*result = *res

	return nil
}
