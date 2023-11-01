package server

import (
	"github.com/abondar24/TaskDisrtibutor/taskApi/handler"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	taskService *service.TaskCommandService
	router      *mux.Router
}

func NewServer(taskService *service.TaskCommandService) *Server {
	apiRouter := mux.NewRouter().StrictSlash(true)
	return &Server{
		taskService,
		apiRouter,
	}
}

func (s *Server) RunServer() {
	createHandler := handler.InitCreateHandler(s.taskService)
	updateHandler := handler.InitUpdateHandler(s.taskService)
	deleteHandler := handler.InitDeleteHandler(s.taskService)

	healthHandler := handler.InitHealthHandler()

	s.router.Methods("POST").Path("/task/create").Handler(createHandler)
	s.router.Methods("PUT").Path("/task/update/{id}").Handler(updateHandler)
	s.router.Methods("DELETE").Path("/task/delete/{id}").Handler(deleteHandler)

	s.router.Methods("GET").Path("/health").Handler(healthHandler)

	http.Handle("/", s.router)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
		return
	}
}
