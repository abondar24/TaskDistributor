package service

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/dao"
)

type TaskSerivce struct {
	taskDAO        *dao.TaskDao
	taskHistoryDAO *dao.TaskHistoryDao
}

type TaskService interface {
	//saveTask also updates UpdatedAt field in case of deletion or updateing
	SaveUpdateTask(task *data.Task) error

	GetTaskById(id *string) *data.Task
	GetTasks(offset *int, limit *int) *[]data.Task
	GetTasksByStatus(status *data.TaskStatus) *[]data.Task
	GetTaskHistory(id *string, offset *int, limit *int) *[]data.Task
}

func NewTaskService(taskDao dao.TaskDao, historyDao dao.TaskHistoryDao) *TaskSerivce {
	return &TaskSerivce{
		taskDAO:        &taskDao,
		taskHistoryDAO: &historyDao,
	}
}

func (ts *TaskSerivce) SaveUpdateTask(task *data.Task) error {
	return nil
}

func (ts *TaskSerivce) GetTaskById(id *string) *data.Task {
	return &data.Task{}
}

func (ts *TaskSerivce) GetTasks(offset *int, limit *int) *[]data.Task {
	return nil
}

func (ts *TaskSerivce) GetTasksByStatus(status *data.TaskStatus) *[]data.Task {
	return nil
}

func (ts *TaskSerivce) GetTaskHistory(id *string, offset *int, limit *int) *[]data.Task {
	return nil
}
