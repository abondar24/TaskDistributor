package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/abondar24/TaskDistributor/taskData/rpc"
	"github.com/abondar24/TaskDistributor/taskReadApi/client"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type RequestHandler struct {
	rpcClient client.Client
}

func NewRequestHandler(rpcClient client.Client) *RequestHandler {
	return &RequestHandler{
		rpcClient: rpcClient,
	}
}

// GetTaskHandler godoc
// @Summary Get task
// @Description Fetch task by id
// @Tags tasks
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} data.Task "Task with latest status"
// @Failure 400 {object} response.ErrorResponse "Wrong id param"
// @Failure 502 {object} response.ErrorResponse "Failed to read from store"
// @Router /task/{id} [get]
func (h *RequestHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handleError(errors.New("wrong ID"), w, http.StatusBadRequest)
		return
	}

	resp, err := h.rpcClient.GetTask(&id)
	if err != nil {
		handleError(err, w, http.StatusBadGateway)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err.Error())
	}
}

// GetTasksByStatusHandler godoc
// @Summary Get tasks by status
// @Description Fetch tasks by specific status with offset and limit
// @Tags tasks
// @Produce  json
// @Param status path string true "Task Status"
// @Param offset query string true "Offset - from which task to fetch"
// @Param limit query string true "Limit - how many tasks to fetch"
// @Success 200 {object} []data.Task "Tasks with specific status"
// @Failure 400 {object} response.ErrorResponse "Wrong path or query param"
// @Failure 502 {object} response.ErrorResponse "Failed to read from store"
// @Router /tasks/{status} [get]
func (h *RequestHandler) GetTasksByStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	status, ok := vars["status"]
	if !ok {
		handleError(errors.New("missing status"), w, http.StatusBadRequest)
		return
	}

	var statusMap = map[string]data.TaskStatus{
		"created":   data.TASK_CREATED,
		"updated":   data.TASK_UPDATED,
		"deleted":   data.TASK_DELETED,
		"completed": data.TASK_COMPLETED,
	}

	taskStatus, err := readStatus(&status, &statusMap)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	offsetStr := r.URL.Query().Get("offset")
	offset, err := readParam(offsetStr)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit, err := readParam(limitStr)
	if err != nil {
		handleError(err, w, http.StatusBadRequest)
		return
	}

	args := rpc.StatusArgs{
		Status: taskStatus,
		Offset: &offset,
		Limit:  &limit,
	}

	resp, err := h.rpcClient.GetTasksByStatus(&args)
	if err != nil {
		handleError(err, w, http.StatusBadGateway)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err.Error())
	}
}

// GetTaskHistoryHandler godoc
// @Summary Get task history
// @Description Fetch task status update history by id
// @Tags tasks
// @Produce  json
// @Param id path string true "Task ID"
// @Success 200 {object} data.TaskHistory "Task with status changes history"
// @Failure 400 {object} response.ErrorResponse "Missing id param"
// @Failure 502 {object} response.ErrorResponse "Failed to read from store"
// @Router /task/history/{id} [get]
func (h *RequestHandler) GetTaskHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		handleError(errors.New("missing ID"), w, http.StatusBadRequest)
		return
	}

	resp, err := h.rpcClient.GetTaskHistory(&id)
	if err != nil {
		handleError(err, w, http.StatusBadGateway)
		return
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err.Error())
	}
}

func handleError(err error, w http.ResponseWriter, errCode int) {
	log.Println(err.Error())
	errorResp := &response.ErrorResponse{
		ERROR: err.Error(),
	}
	w.WriteHeader(errCode)

	err = json.NewEncoder(w).Encode(errorResp)
	if err != nil {
		log.Println(err.Error())
	}
}

func readParam(param string) (int, error) {
	if param == "" {
		return -1, errors.New("missing parameter")
	}

	res, err := strconv.Atoi(param)
	if err != nil {
		return -1, errors.New("wrong parameter")
	}

	return res, nil
}

func readStatus(status *string, statusMap *map[string]data.TaskStatus) (*data.TaskStatus, error) {
	if statusMap == nil {
		return nil, fmt.Errorf("statusMap is nil")
	}

	if res, ok := (*statusMap)[*status]; ok {
		return &res, nil
	}

	return nil, fmt.Errorf("invalid status value: %s", *status)
}
