// Code generated by MockGen. DO NOT EDIT.
// Source: service/queueService.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/abondar24/TaskDisrtibutor/taskApi/model"
	gomock "github.com/golang/mock/gomock"
)

// MockQueueService is a mock of QueueService interface.
type MockQueueService struct {
	ctrl     *gomock.Controller
	recorder *MockQueueServiceMockRecorder
}

// MockQueueServiceMockRecorder is the mock recorder for MockQueueService.
type MockQueueServiceMockRecorder struct {
	mock *MockQueueService
}

// NewMockQueueService creates a new mock instance.
func NewMockQueueService(ctrl *gomock.Controller) *MockQueueService {
	mock := &MockQueueService{ctrl: ctrl}
	mock.recorder = &MockQueueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueService) EXPECT() *MockQueueServiceMockRecorder {
	return m.recorder
}

// PublishToQueue mocks base method.
func (m *MockQueueService) PublishToQueue(task *model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishToQueue", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishToQueue indicates an expected call of PublishToQueue.
func (mr *MockQueueServiceMockRecorder) PublishToQueue(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishToQueue", reflect.TypeOf((*MockQueueService)(nil).PublishToQueue), task)
}
