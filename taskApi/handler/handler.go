package handler

import (
	"encoding/json"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type RequestHandler struct {
	taskService *service.TaskCommandService
}

func NewHandler(taskService *service.TaskCommandService) *RequestHandler {
	return &RequestHandler{
		taskService: taskService,
	}
}

// CreateTaskHandler godoc
// @Summary Create a new task
// @Description Create a new task and send it to the queue
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body model.TaskRequest true "Create Task"
// @Success 200 {object} model.TaskResponse
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/create [post]
func (h *RequestHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task model.TaskRequest
	err := json.NewDecoder(r.Body).Decode(&task)

	id, err := h.taskService.CreateTask(&task.Name)
	if err != nil {
		handleError(err, w)
	}

	json.NewEncoder(w).Encode(model.TaskResponse{
		ID: id,
	})
}

// UpdateTaskHandler godoc
// @Summary Update task
// @Description Change status of existing task
// @Tags tasks
// @Produce  json
// @Param id path string true "Update Task"
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/update [put]
func (h *RequestHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.taskService.UpdateTask(&id)
	if err != nil {
		handleError(err, w)
	}

}

// DeleteTaskHandler godoc
// @Summary Delete task
// @Description Delete existing task
// @Tags tasks
// @Produce  json
// @Param id path string true "Delete Task"
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/delete [delete]
func (h *RequestHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.taskService.DeleteTask(&id)
	if err != nil {
		handleError(err, w)
	}

}

// HealthHandler godoc
// @Summary Health of service
// @Description Checks if service is up
// @Tags tasks
// @Produce  json
// @Success 200 {object} model.HealthResponse
// @Router /health [get]
func (h *RequestHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(model.HealthResponse{MESSAGE: "TaskAPI is up"})

}

func handleError(err error, w http.ResponseWriter) {
	log.Println(err.Error())
	errorResp := &model.ErrorResponse{
		ERROR: err.Error(),
	}
	json.NewEncoder(w).Encode(errorResp)
}
