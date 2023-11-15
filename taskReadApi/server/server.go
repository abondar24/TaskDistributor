package server

import (
	jsonparse "encoding/json"
	"github.com/abondar24/TaskDistributor/taskData/response"
	_ "github.com/abondar24/TaskDistributor/taskReadApi/docs"
	"github.com/abondar24/TaskDistributor/taskReadApi/handler"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	requestHandler *handler.RequestHandler
	router         *mux.Router
}

func NewServer(requestHandler *handler.RequestHandler) *Server {
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		requestHandler,
		router,
	}
}

// @title Task Read API
// @version 1.0
// @description Task API to read tasks from store
// @termsOfService http://swagger.io/terms/
// @contact.name Alex
// @contact.email abondar24@yahoo.com
// @license.name MIT
// @host localhost:8082
// @BasePath /

func (srv *Server) RunServer(port int) {
	srv.router.HandleFunc("/health", healthHandler).Methods("GET")
	srv.router.HandleFunc("/task/{id}", srv.requestHandler.GetTaskHandler).Methods("GET")
	srv.router.HandleFunc("/tasks/{status}", srv.requestHandler.GetTasksByStatusHandler).Methods("GET")
	srv.router.HandleFunc("/task/history/{id}", srv.requestHandler.GetTaskHistoryHandler).Methods("GET")
	srv.router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.Handle("/", srv.router)
	err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := jsonparse.NewEncoder(w).Encode(response.HealthResponse{MESSAGE: "Task Read API is up"})
	if err != nil {
		log.Fatalln(err.Error())
	}

}
