package main

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/handler"
	"github.com/abondar24/TaskDisrtibutor/taskApi/server"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
)

func main() {
	amqpService := service.NewAmqpService()
	taskService := service.NewTaskService(amqpService)
	requestHandler := handler.NewHandler(taskService)

	srv := server.NewServer(requestHandler)
	srv.RunServer()
}
