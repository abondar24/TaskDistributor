package server

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/handler"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	taskService *service.TaskService
	router      *mux.Router
}

func NewServer(taskService *service.TaskService) *Server {
	apiRouter := mux.NewRouter().StrictSlash(true)
	return &Server{
		taskService,
		apiRouter,
	}
}

func (s *Server) RunServer() {
	taskHandler := handler.InitSendHandler(s.taskService)
	healthHandler := handler.InitHealthHandler()

	s.router.Methods("POST").Path("/tasks/send").Handler(taskHandler)
	s.router.Methods("GET").Path("/health").Handler(healthHandler)

	http.Handle("/", s.router)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
		return
	}
}
