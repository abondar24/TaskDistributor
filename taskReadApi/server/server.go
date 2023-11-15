package server

import (
	jsonparse "encoding/json"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		router,
	}
}

func (srv *Server) RunServer(port int) {
	srv.router.HandleFunc("/health", healthHandler).Methods("GET")

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
