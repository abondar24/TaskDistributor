package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/abondar24/TaskDistributor/taskApi/model"
	"github.com/abondar24/TaskDistributor/taskApi/service"
	"github.com/abondar24/TaskDistributor/taskData/response"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler_CreateTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskService := service.NewMockTaskService(ctrl)
	test := "test"
	taskService.EXPECT().CreateTask(&test).Return("test", nil)

	requestHandler := NewRequestHandler(taskService)

	task := model.TaskRequest{Name: "test"}
	taskJson, err := json.Marshal(task)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(taskJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var taskResp model.TaskResponse
	err = json.Unmarshal(rr.Body.Bytes(), &taskResp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, test, taskResp.ID)
}

func TestRequestHandler_CreateTaskHandlerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskService := service.NewMockTaskService(ctrl)
	test := "test"
	testErr := errors.New("test error")
	taskService.EXPECT().CreateTask(&test).Return("test", testErr)

	requestHandler := NewRequestHandler(taskService)

	task := model.TaskRequest{Name: "test"}
	taskJson, err := json.Marshal(task)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/task", bytes.NewBuffer(taskJson))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.CreateTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadGateway)
	}

	var errorResp model.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testErr.Error(), errorResp.ERROR)
}

func TestRequestHandler_UpdateTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskService := service.NewMockTaskService(ctrl)
	test := "123"
	complete := false
	taskService.EXPECT().UpdateTask(&test, &complete).Return(nil)

	requestHandler := NewRequestHandler(taskService)

	req, err := http.NewRequest("PUT", "/task/123?complete=false", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "123"}) // Set the path parameter "id"

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.UpdateTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

}

func TestRequestHandler_UpdateTaskHandlerIdMissing(t *testing.T) {

	requestHandler := NewRequestHandler(nil)

	req, err := http.NewRequest("PUT", "/task", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.UpdateTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

	var errorResp model.ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &errorResp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "missing ID", errorResp.ERROR)
}

func TestRequestHandler_UpdateTaskHandlerNoComplete(t *testing.T) {

	requestHandler := NewRequestHandler(nil)

	req, err := http.NewRequest("PUT", "/task/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.UpdateTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

}

func TestRequestHandler_UpdateTaskHandlerWrongComplete(t *testing.T) {

	requestHandler := NewRequestHandler(nil)

	req, err := http.NewRequest("PUT", "/task/123?complete=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.UpdateTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
	}

}

func TestRequestHandler_DeleteTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskService := service.NewMockTaskService(ctrl)
	test := "123"
	taskService.EXPECT().DeleteTask(&test).Return(nil)

	requestHandler := NewRequestHandler(taskService)

	req, err := http.NewRequest("DELETE", "/task/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "123"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.DeleteTaskHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

}

func TestRequestHandler_HealthHandler(t *testing.T) {

	requestHandler := NewRequestHandler(nil)

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(requestHandler.HealthHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var healthResp response.HealthResponse
	err = json.Unmarshal(rr.Body.Bytes(), &healthResp)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "TaskAPI is up", healthResp.MESSAGE)

}
