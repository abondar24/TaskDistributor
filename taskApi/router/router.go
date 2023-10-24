package router

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func InitRouter() Router {
	taskRouter := mux.NewRouter()

	return Router{router: taskRouter}
}

func (r *Router) AddTaskRoute(handler *httptransport.Server) {
	r.router.Methods("POST").Path("/send").Handler(handler)
}
