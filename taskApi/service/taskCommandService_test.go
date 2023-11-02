package service

import (
	"errors"
	"github.com/abondar24/TaskDistributor/taskApi/queue"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskCommandService_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpService := queue.NewMockProducer(ctrl)
	amqpService.EXPECT().PublishToQueue(gomock.Any()).Return(nil)

	taskService := NewTaskService(amqpService)

	task := "test"
	result, err := taskService.CreateTask(&task)

	assert.Nil(t, err)
	assert.NotNil(t, result)

}

func TestTaskCommandService_CreateTaskError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	err := errors.New("some error")
	amqpService := queue.NewMockProducer(ctrl)
	amqpService.EXPECT().PublishToQueue(gomock.Any()).Return(err)

	taskService := NewTaskService(amqpService)

	task := "test"
	result, err := taskService.CreateTask(&task)

	assert.NotNil(t, err)
	assert.Equal(t, "some error", err.Error())
	assert.NotNil(t, result)

}

func TestTaskCommandService_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpService := queue.NewMockProducer(ctrl)
	amqpService.EXPECT().PublishToQueue(gomock.Any()).Return(nil)

	taskService := NewTaskService(amqpService)

	task := "test"
	completed := false

	err := taskService.UpdateTask(&task, &completed)
	assert.Nil(t, err)
}

func TestTaskCommandService_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpService := queue.NewMockProducer(ctrl)
	amqpService.EXPECT().PublishToQueue(gomock.Any()).Return(nil)

	taskService := NewTaskService(amqpService)

	task := "test"

	err := taskService.DeleteTask(&task)
	assert.Nil(t, err)
}
