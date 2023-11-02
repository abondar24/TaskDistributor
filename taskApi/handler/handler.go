package handler

import (
	"encoding/json"
	"errors"
	"github.com/abondar24/TaskDistributor/taskApi/model"
	"github.com/abondar24/TaskDistributor/taskApi/service"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type RequestHandler struct {
	taskService service.TaskService
}

func NewRequestHandler(taskService service.TaskService) *RequestHandler {
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
// @BadGateway 502 {object} model.ErrorResponse
// @Router /task [post]
func (h *RequestHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task model.TaskRequest
	err := json.NewDecoder(r.Body).Decode(&task)

	id, err := h.taskService.CreateTask(&task.Name)
	if err != nil {
		handleError(err, w, http.StatusBadGateway)
		return
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
// @Param complete query string true "Complete task. Possible values: true/false"
// @BadRequest 400 {object} model.ErrorResponse
// @BadGateway 502 {object} model.ErrorResponse
// @Router /task/{id} [put]
func (h *RequestHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handleError(errors.New("missing ID"), w, http.StatusBadRequest)
		return
	}

	completeStr := r.URL.Query().Get("complete")
	if completeStr == "" {
		handleError(errors.New("missing 'complete' parameter"), w, http.StatusBadRequest)
		return
	}

	complete, err := strconv.ParseBool(completeStr)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	err = h.taskService.UpdateTask(&id, &complete)
	if err != nil {
		handleError(err, w, http.StatusBadGateway)
		return
	}

}

// DeleteTaskHandler godoc
// @Summary Delete task
// @Description Delete existing task
// @Tags tasks
// @Produce  json
// @Param id path string true "Task ID"
// @BadRequest 400 {object} model.ErrorResponse
// @BadGateway 502 {object} model.ErrorResponse
// @Router /task/{id} [delete]
func (h *RequestHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := h.taskService.DeleteTask(&id)
	if err != nil {
		handleError(err, w, http.StatusBadGateway)
	}

}

func (h *RequestHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response.HealthResponse{MESSAGE: "TaskAPI is up"})

}

func handleError(err error, w http.ResponseWriter, errCode int) {
	log.Println(err.Error())
	errorResp := &model.ErrorResponse{
		ERROR: err.Error(),
	}
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(errorResp)
}
