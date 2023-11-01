package server

import (
	_ "github.com/abondar24/TaskDisrtibutor/taskApi/docs"
	"github.com/abondar24/TaskDisrtibutor/taskApi/handler"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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

// @title Task API
// @version 1.0
// @description Task API to send commands as tasks
// @termsOfService http://swagger.io/terms/
// @contact.name Alex
// @contact.email abondar24@yahoo.com
// @license.name MIT
// @host localhost:8080
// @BasePath /

func (s *Server) RunServer() {
	createHandler := handler.InitCreateHandler(s.taskService)
	updateHandler := handler.InitUpdateHandler(s.taskService)
	deleteHandler := handler.InitDeleteHandler(s.taskService)

	healthHandler := handler.InitHealthHandler()

	s.router.Methods("POST").Path("/task/create").Handler(createHandler)
	s.router.Methods("PUT").Path("/task/update/{id}").Handler(updateHandler)
	s.router.Methods("DELETE").Path("/task/delete/{id}").Handler(deleteHandler)

	s.router.Methods("GET").Path("/health").Handler(healthHandler)

	//todo - add endpoint for completeing task(update can be extended too)

	s.router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.Handle("/", s.router)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
		return
	}
}
