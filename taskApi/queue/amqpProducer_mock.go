// Code generated by MockGen. DO NOT EDIT.
// Source: queue/queueProducer.go

// Package queue is a generated GoMock package.
package queue

import (
	reflect "reflect"

	model "github.com/abondar24/TaskDistributor/taskApi/model"
	gomock "github.com/golang/mock/gomock"
)

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// PublishToQueue mocks base method.
func (m *MockProducer) PublishToQueue(task *model.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishToQueue", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishToQueue indicates an expected call of PublishToQueue.
func (mr *MockProducerMockRecorder) PublishToQueue(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishToQueue", reflect.TypeOf((*MockProducer)(nil).PublishToQueue), task)
}
