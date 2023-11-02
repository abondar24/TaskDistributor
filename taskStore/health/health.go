package health

import (
	"encoding/json"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/gorilla/mux"
	"net/http"
)

type Health struct {
	router *mux.Router
}

func NewHealth() *Health {
	healthRouter := mux.NewRouter().StrictSlash(true)
	return &Health{
		healthRouter,
	}
}

func (h *Health) RunServer() {

	h.router.HandleFunc("/health", healthHandler).Methods("GET")

	http.Handle("/", h.router)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		panic(err)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response.HealthResponse{MESSAGE: "Task Store is up"})

}
