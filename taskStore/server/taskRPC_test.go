package server

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskData/rpc"
	"github.com/abondar24/TaskDistributor/taskStore/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskRPC_GetTask(t *testing.T) {
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

	taskService := service.NewMockTaskService(ctrl)
	taskService.EXPECT().GetTaskById(&id).Return(&task, nil)

	taskRPC := NewTaskRPC(taskService)

	resp := &data.Task{}
	err := taskRPC.GetTask(nil, &id, resp)

	assert.Nil(t, err)
	assert.Equal(t, task.ID, resp.ID)
	assert.Equal(t, task.Name, resp.Name)
	assert.Equal(t, task.Status, resp.Status)
	assert.Equal(t, task.CreateTime, resp.CreateTime)
	assert.Equal(t, task.UpdateTime, resp.UpdateTime)
}

func TestTaskRPC_GetTaskByStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	offset := 0
	limit := 1
	taskStatus := data.TASK_CREATED

	queryArgs := rpc.StatusArgs{
		Status: &taskStatus,
		Offset: &offset,
		Limit:  &limit,
	}

	id := "test"
	task := data.Task{
		ID:         id,
		Name:       "test",
		Status:     data.TASK_CREATED,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	tasks := make([]*data.Task, 0)
	tasks = append(tasks, &task)

	taskService := service.NewMockTaskService(ctrl)
	taskService.EXPECT().GetTasksByStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(tasks, nil)

	taskRPC := NewTaskRPC(taskService)

	resp := make([]*data.Task, 0)
	err := taskRPC.GetTasksByStatus(nil, &queryArgs, &resp)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, task.ID, resp[0].ID)
	assert.Equal(t, task.Name, resp[0].Name)
	assert.Equal(t, task.Status, resp[0].Status)
	assert.Equal(t, task.CreateTime, resp[0].CreateTime)
	assert.Equal(t, task.UpdateTime, resp[0].UpdateTime)
}

func TestTaskRPC_GetTaskHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"

	statusHistory := make([]data.TaskStatusEntry, 0)
	taskHistoryEntry := data.TaskStatusEntry{
		Status:    data.TASK_CREATED,
		UpdatedAt: time.Now(),
	}

	statusHistory = append(statusHistory, taskHistoryEntry)

	taskHistory := data.TaskHistory{
		ID:            id,
		Name:          "test",
		CreateTime:    time.Now(),
		StatusHistory: statusHistory,
	}

	taskService := service.NewMockTaskService(ctrl)
	taskService.EXPECT().GetTaskHistory(&id).Return(&taskHistory, nil)

	taskRPC := NewTaskRPC(taskService)

	resp := &data.TaskHistory{}
	err := taskRPC.GetTaskHistory(nil, &id, resp)

	assert.Nil(t, err)
	assert.Equal(t, taskHistory.ID, resp.ID)
	assert.Equal(t, taskHistory.Name, resp.Name)
	assert.Equal(t, len(taskHistory.StatusHistory), len(resp.StatusHistory))
	assert.Equal(t, taskHistory.StatusHistory[0].Status, resp.StatusHistory[0].Status)
	assert.Equal(t, taskHistory.StatusHistory[0].UpdatedAt, resp.StatusHistory[0].UpdatedAt)

}
