package server

import (
	jsonparse "encoding/json"
	"github.com/abondar24/TaskDistributor/taskData/response"
	_ "github.com/abondar24/TaskDistributor/taskStore/docs"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type Server struct {
	router  *mux.Router
	taskRPC *TaskRPC
}

func NewServer(taskRPC *TaskRPC) *Server {
	router := mux.NewRouter().StrictSlash(true)
	return &Server{
		router,
		taskRPC,
	}
}

// RunServer @title Task Store
// @version 1.0
// @description Task store - accepts commands and exposes JSON-RPC API
// @termsOfService http://swagger.io/terms/
// @contact.name Alex
// @contact.email abondar24@yahoo.com
// @license.name MIT
// @host localhost:8081
// @BasePath /
func (srv *Server) RunServer(port string) {

	rpcSrv := rpc.NewServer()
	rpcSrv.RegisterCodec(json.NewCodec(), "application/json")
	err := rpcSrv.RegisterService(srv.taskRPC, "TaskRPC")
	if err != nil {
		log.Fatalln(err.Error())
	}

	srv.router.HandleFunc("/health", healthHandler).Methods("GET")
	srv.router.Handle("/rpc", rpcSrv).Methods("POST")
	srv.router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	http.Handle("/", srv.router)
	err = http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		panic(err)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := jsonparse.NewEncoder(w).Encode(response.HealthResponse{MESSAGE: "Task Store is up"})
	if err != nil {
		log.Fatalln(err.Error())
	}

}
