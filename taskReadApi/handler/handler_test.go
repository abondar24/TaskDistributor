package handler

import (
	"encoding/json"
	"errors"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/abondar24/TaskDistributor/taskData/rpc"
	"github.com/abondar24/TaskDistributor/taskReadApi/client"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestHandler_GetTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"
	task := data.Task{
		ID:         id,
		Name:       "test",
		Status:     data.TASK_CREATED,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	taskRpcClient := client.NewMockClient(ctrl)
	taskRpcClient.EXPECT().GetTask(&id).Return(&task, nil)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/task/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": id}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var resp data.Task
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, task.ID, resp.ID)
	assert.Equal(t, task.Name, resp.Name)
}

func TestRequestHandler_GetTaskHandlerBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRpcClient := client.NewMockClient(ctrl)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/task", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "wrong ID", resp.ERROR)
}

func TestRequestHandler_GetTaskHandlerBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"
	testErr := errors.New("test err")

	taskRpcClient := client.NewMockClient(ctrl)
	taskRpcClient.EXPECT().GetTask(&id).Return(nil, testErr)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/task/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": id}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadGateway)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testErr.Error(), resp.ERROR)
}

func TestRequestHandler_GetTasksByStatusHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := data.Task{
		ID:         "test",
		Name:       "test",
		Status:     data.TASK_CREATED,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	tasks := []*data.Task{&task}

	stat := data.TASK_CREATED
	offset := 0
	limit := 1
	args := rpc.StatusArgs{
		Status: &stat,
		Offset: &offset,
		Limit:  &limit,
	}

	taskRpcClient := client.NewMockClient(ctrl)
	taskRpcClient.EXPECT().GetTasksByStatus(&args).Return(&tasks, nil)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/tasks/created?offset=0&limit=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"status": stat}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTasksByStatusHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var resp []*data.Task
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(resp))
	assert.Equal(t, task.ID, resp[0].ID)
}

func TestRequestHandler_GetTasksByStatusHandlerMissingStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRpcClient := client.NewMockClient(ctrl)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTasksByStatusHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "missing status", resp.ERROR)
}

func TestRequestHandler_GetTasksByStatusHandlerWrongStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRpcClient := client.NewMockClient(ctrl)
	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/tasks/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"status": "test"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTasksByStatusHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "invalid status value: test", resp.ERROR)
}

func TestRequestHandler_GetTasksByStatusHandlerMissingParameter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRpcClient := client.NewMockClient(ctrl)
	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/tasks/updated?offset=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"status": "updated"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTasksByStatusHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "missing parameter", resp.ERROR)
}

func TestRequestHandler_GetTasksByStatusHandlerWrongParameter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRpcClient := client.NewMockClient(ctrl)
	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/tasks/updated?offset=1&limit=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"status": "updated"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTasksByStatusHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "wrong parameter", resp.ERROR)
}

func TestRequestHandler_GetTasksByStatusHandlerBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stat := data.TASK_CREATED
	offset := 0
	limit := 1
	args := rpc.StatusArgs{
		Status: &stat,
		Offset: &offset,
		Limit:  &limit,
	}

	testErr := errors.New("test err")

	taskRpcClient := client.NewMockClient(ctrl)
	taskRpcClient.EXPECT().GetTasksByStatus(&args).Return(nil, testErr)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/tasks/created?offset=0&limit=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"status": stat}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTasksByStatusHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadGateway)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testErr.Error(), resp.ERROR)
}

func TestRequestHandler_GetTaskHistoryHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"

	statusEntry := data.TaskStatusEntry{
		Status:    data.TASK_CREATED,
		UpdatedAt: time.Now(),
	}

	statusEntries := []data.TaskStatusEntry{statusEntry}

	taskHistory := data.TaskHistory{
		ID:            id,
		Name:          "test",
		CreateTime:    time.Now(),
		StatusHistory: statusEntries,
	}

	taskRpcClient := client.NewMockClient(ctrl)
	taskRpcClient.EXPECT().GetTaskHistory(&id).Return(&taskHistory, nil)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/task/history/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": id}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTaskHistoryHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var resp data.Task
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, taskHistory.ID, resp.ID)
	assert.Equal(t, taskHistory.Name, resp.Name)
	assert.Equal(t, 1, len(taskHistory.StatusHistory))
	assert.Equal(t, statusEntry.Status, taskHistory.StatusHistory[0].Status)
}

func TestRequestHandler_GetTaskHistoryHandlerBadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRpcClient := client.NewMockClient(ctrl)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/task/history", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTaskHistoryHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "missing ID", resp.ERROR)
}

func TestRequestHandler_GetTaskHistoryHandlerBadGateway(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"

	testErr := errors.New("test err")

	taskRpcClient := client.NewMockClient(ctrl)
	taskRpcClient.EXPECT().GetTaskHistory(&id).Return(nil, testErr)

	requestHandler := NewRequestHandler(taskRpcClient)

	req, err := http.NewRequest("GET", "/task/history/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": id}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.GetTaskHistoryHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadGateway)
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testErr.Error(), resp.ERROR)
}
