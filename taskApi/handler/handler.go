package handler

import (
	"encoding/json"
	"github.com/abondar24/TaskDisrtibutor/taskApi/model"
	"github.com/abondar24/TaskDisrtibutor/taskApi/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
// @Param task body model.TaskRequest  true "Task name"
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
// @Param id path string true "Task ID"
// @Param complete query string true "Complete task"
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/update/{id} [put]
func (h *RequestHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	complete, err := strconv.ParseBool(r.URL.Query()["complete"][0])
	if err != nil {
		handleError(err, w)
	}

	err = h.taskService.UpdateTask(&id, &complete)
	if err != nil {
		handleError(err, w)
	}

}

// DeleteTaskHandler godoc
// @Summary Delete task
// @Description Delete existing task
// @Tags tasks
// @Produce  json
// @Param id path string true "Task ID"
// @BadRequest 400 {object} model.ErrorResponse
// @Router /task/delete/{id} [delete]
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
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(errorResp)
}
