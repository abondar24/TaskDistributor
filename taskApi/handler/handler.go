package handler

import (
	"context"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func InitCreateHandler(ts *service.TaskCommandService) *httptransport.Server {
	return httptransport.NewServer(
		initCreateEndpoint(ts),
		decodeTaskRequest,
		encodeResponse)

}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task and send it to the queue
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body model.TaskRequest true "Create Task"
// @Success 200 {object} model.TaskResponse
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/create [post]
func initCreateEndpoint(ts *service.TaskCommandService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TaskRequest)

		id, err := ts.CreateTask(&req.Name)
		if err != nil {
			return model.ErrorResponse{
				ERROR: err.Error(),
			}, nil
		}

		return model.TaskResponse{ID: id}, nil
	}
}

func InitUpdateHandler(ts *service.TaskCommandService) *httptransport.Server {
	return httptransport.NewServer(
		initUpdateEndpoint(ts),
		decodeTaskRequest,
		encodeResponse)

}

// UpdateTask godoc
// @Summary Update task
// @Description Change status of existing task
// @Tags tasks
// @Produce  json
// @Param id path string true "Update Task"
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/update [put]
func initUpdateEndpoint(ts *service.TaskCommandService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TaskRequest)
		err := ts.UpdateTask(&req.ID)
		if err != nil {
			return model.ErrorResponse{
				ERROR: err.Error(),
			}, nil
		}

		return nil, nil
	}
}

func InitDeleteHandler(ts *service.TaskCommandService) *httptransport.Server {
	return httptransport.NewServer(
		initDeleteEndpoint(ts),
		decodeTaskRequest,
		encodeResponse)

}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete existing task
// @Tags tasks
// @Produce  json
// @Param id path string true "Delete Task"
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/delete [delete]
func initDeleteEndpoint(ts *service.TaskCommandService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TaskRequest)
		err := ts.DeleteTask(&req.ID)
		if err != nil {
			return model.ErrorResponse{
				ERROR: err.Error(),
			}, nil
		}

		return nil, nil
	}
}

func InitHealthHandler() *httptransport.Server {
	return httptransport.NewServer(
		initHealthEndpoint(),
		decodeHealthRequest,
		encodeResponse)

}

// Healthcheck godoc
// @Summary Health of service
// @Description Checks if service is up
// @Tags tasks
// @Produce  json
// @Success 200 {object} model.HealthResponse
// @Router /health [get]
func initHealthEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return model.HealthResponse{MESSAGE: "TaskAPI is up"}, nil
	}
}
