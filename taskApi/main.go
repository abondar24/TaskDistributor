package main

import (
	"github.com/abondar24/TaskDistributor/taskApi/handler"
	"github.com/abondar24/TaskDistributor/taskApi/server"
	"github.com/abondar24/TaskDistributor/taskApi/service"
)

func main() {
	amqpService := service.NewAmqpService()
	taskService := service.NewTaskService(amqpService)
	requestHandler := handler.NewRequestHandler(taskService)

	srv := server.NewServer(requestHandler)
	srv.RunServer()
}
