package server

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
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

/// r parameter is needed by gorilla rpvc in order to register service

func (tr *TaskRPC) GetTask(r *http.Request, id *string, response *data.Task) error {

	log.Printf("Fetching task by id %s\n", *id)

	res, err := tr.taskService.GetTaskById(id)
	if err != nil {
		return err
	}

	*response = *res

	log.Println("Got response")
	log.Println(response)

	return nil
}

// TODO fix method not found
func (tr *TaskRPC) GetTasksByStatus(r *http.Request, status *data.TaskStatus, offset *int, limit *int, response *[]*data.Task) error {
	log.Printf("Fetching tasks by status %s\n", *status)

	res, err := tr.taskService.GetTasksByStatus(status, offset, limit)
	if err != nil {
		return err
	}

	*response = res

	log.Println("Got response")
	log.Println(response)

	return nil
}

func (tr *TaskRPC) GetTaskHistory(r *http.Request, id *string, response *data.TaskHistory) error {
	log.Printf("Fetching task history for id %s\n", *id)

	res, err := tr.taskService.GetTaskHistory(id)
	if err != nil {
		return err
	}

	*response = *res

	log.Println("Got response")
	log.Println(response)

	return nil
}
