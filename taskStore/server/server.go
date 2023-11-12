package server

import (
	"encoding/json"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct {
	router  *mux.Router
	taskRPC *TaskRPC
}

func NewServer(taskRPC *TaskRPC) *Server {
	healthRouter := mux.NewRouter().StrictSlash(true)
	return &Server{
		healthRouter,
		taskRPC,
	}
}

func (srv *Server) RunServer(port string) {

	err := rpc.Register(srv.taskRPC)
	if err != nil {
		log.Fatalln(err.Error())
	}

	srv.router.HandleFunc("/server", healthHandler).Methods("GET")
	srv.router.HandleFunc("/rpc", rpcHandler).Methods("POST")

	http.Handle("/", srv.router)
	err = http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		panic(err)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response.HealthResponse{MESSAGE: "Task Store is up"})

}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}

	conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "Failed to hijack connection", http.StatusInternalServerError)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}(conn)

	jsonrpc.ServeConn(conn)
}
