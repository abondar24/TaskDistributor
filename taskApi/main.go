package main

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/handler"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//TODO add rabbitMQ
	taskService := &service.TaskService{}

	taskHandler := handler.InitSendHandler(taskService)
	healthHandler := handler.InitHealthHandler()

	apiRouter := mux.NewRouter()

	apiRouter.Methods("POST").Path("/tasks/send").Handler(taskHandler)
	apiRouter.Methods("GET").Path("/health").Handler(healthHandler)

	http.Handle("/", apiRouter)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
		return
	}
}
