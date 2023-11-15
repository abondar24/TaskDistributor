package server

import (
	_ "github.com/abondar24/TaskDistributor/taskApi/docs"
	"github.com/abondar24/TaskDistributor/taskApi/handler"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"strconv"
)

type Server struct {
	requestHandler *handler.RequestHandler
	router         *mux.Router
}

func NewServer(requestHandler *handler.RequestHandler) *Server {
	apiRouter := mux.NewRouter().StrictSlash(true)
	return &Server{
		requestHandler,
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

func (s *Server) RunServer(port int) {

	s.router.HandleFunc("/task", s.requestHandler.CreateTaskHandler).Methods("POST")
	s.router.HandleFunc("/task/{id}", s.requestHandler.UpdateTaskHandler).Methods("PUT")
	s.router.HandleFunc("/task/{id}", s.requestHandler.DeleteTaskHandler).Methods("DELETE")
	s.router.HandleFunc("/health", s.requestHandler.HealthHandler).Methods("GET")
	s.router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.Handle("/", s.router)
	err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
		return
	}
}
