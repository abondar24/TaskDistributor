package main

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/server"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
)

func main() {
	//TODO add rabbitMQ
	taskService := &service.TaskService{}

	srv := server.NewServer(taskService)
	srv.RunServer()
}
