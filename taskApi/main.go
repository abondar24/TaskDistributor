package main

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/server"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
)

func main() {
	amqpService := service.NewAmqpService()
	taskService := service.NewTaskService(amqpService)

	srv := server.NewServer(taskService)
	srv.RunServer()
}
