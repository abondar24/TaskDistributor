package handler

import (
	"context"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func InitCreateHandler(ts *service.TaskService) *httptransport.Server {
	return httptransport.NewServer(
		initCreateEndpoint(ts),
		decodeTaskRequest,
		encodeResponse)

}

func initCreateEndpoint(ts *service.TaskService) endpoint.Endpoint {
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

func InitUpdateHandler(ts *service.TaskService) *httptransport.Server {
	return httptransport.NewServer(
		initUpdateEndpoint(ts),
		decodeTaskRequest,
		encodeResponse)

}

func initUpdateEndpoint(ts *service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TaskRequest)
		err := ts.UpdateTask(&req.Name, &req.ID)
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

func initHealthEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return model.HealthResponse{MESSAGE: "TaskAPI is up"}, nil
	}
}
