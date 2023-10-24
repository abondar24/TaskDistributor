package handler

import (
	"context"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func InitSendHandler(ts *service.TaskService) *httptransport.Server {
	return httptransport.NewServer(
		initSendEndpoint(ts),
		decodeTaskRequest,
		encodeTaskResponse)

}

func initSendEndpoint(ts *service.TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.TypeRequest)

		id, err := ts.SendTask(req.Name, req.Status)
		if err != nil {
			return model.ErrorResponse{
				ERROR: err.Error(),
			}, nil
		}

		return model.TaskResponse{ID: id}, nil
	}
}
