// Code generated by MockGen. DO NOT EDIT.
// Source: service/taskService.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	data "github.com/abondar24/TaskDistributor/taskData/data"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskService is a mock of TaskService interface.
type MockTaskService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceMockRecorder
}

// MockTaskServiceMockRecorder is the mock recorder for MockTaskService.
type MockTaskServiceMockRecorder struct {
	mock *MockTaskService
}

// NewMockTaskService creates a new mock instance.
func NewMockTaskService(ctrl *gomock.Controller) *MockTaskService {
	mock := &MockTaskService{ctrl: ctrl}
	mock.recorder = &MockTaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskService) EXPECT() *MockTaskServiceMockRecorder {
	return m.recorder
}

// GetTaskById mocks base method.
func (m *MockTaskService) GetTaskById(id *string) (*data.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskById", id)
	ret0, _ := ret[0].(*data.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskById indicates an expected call of GetTaskById.
func (mr *MockTaskServiceMockRecorder) GetTaskById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskById", reflect.TypeOf((*MockTaskService)(nil).GetTaskById), id)
}

// GetTaskHistory mocks base method.
func (m *MockTaskService) GetTaskHistory(id *string) (*data.TaskHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskHistory", id)
	ret0, _ := ret[0].(*data.TaskHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskHistory indicates an expected call of GetTaskHistory.
func (mr *MockTaskServiceMockRecorder) GetTaskHistory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskHistory", reflect.TypeOf((*MockTaskService)(nil).GetTaskHistory), id)
}

// GetTasksByStatus mocks base method.
func (m *MockTaskService) GetTasksByStatus(status *data.TaskStatus, offset, limit *int) ([]*data.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByStatus", status, offset, limit)
	ret0, _ := ret[0].([]*data.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByStatus indicates an expected call of GetTasksByStatus.
func (mr *MockTaskServiceMockRecorder) GetTasksByStatus(status, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByStatus", reflect.TypeOf((*MockTaskService)(nil).GetTasksByStatus), status, offset, limit)
}

// SaveUpdateTask mocks base method.
func (m *MockTaskService) SaveUpdateTask(task *data.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUpdateTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUpdateTask indicates an expected call of SaveUpdateTask.
func (mr *MockTaskServiceMockRecorder) SaveUpdateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUpdateTask", reflect.TypeOf((*MockTaskService)(nil).SaveUpdateTask), task)
}
