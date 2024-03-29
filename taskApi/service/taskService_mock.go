// Code generated by MockGen. DO NOT EDIT.
// Source: service/taskService.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

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

// CreateTask mocks base method.
func (m *MockTaskService) CreateTask(name *string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", name)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskServiceMockRecorder) CreateTask(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskService)(nil).CreateTask), name)
}

// DeleteTask mocks base method.
func (m *MockTaskService) DeleteTask(id *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskServiceMockRecorder) DeleteTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskService)(nil).DeleteTask), id)
}

// UpdateTask mocks base method.
func (m *MockTaskService) UpdateTask(id *string, completed *bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", id, completed)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskServiceMockRecorder) UpdateTask(id, completed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskService)(nil).UpdateTask), id, completed)
}
