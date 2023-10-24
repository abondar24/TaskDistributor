package main

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/handler"
	"github.com/abondar24/TaskDisrtibutor/taskApi/router"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"net/http"
)

func main() {
	//TODO add rabbitMQ
	taskService := &service.TaskService{}

	taskHandler := handler.InitSendHandler(taskService)
	taskRouter := router.InitRouter()
	taskRouter.AddTaskRoute(taskHandler)

	http.Handle("/tasks", taskHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		return
	}
}
