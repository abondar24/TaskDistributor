package service

import (
	"database/sql"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/dao"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskService_SaveTask(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := data.Task{
		ID:         "test",
		Name:       "test",
		Status:     data.TASK_CREATED,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	dbService := NewMockStoreService(ctrl)
	dbService.EXPECT().BeginTx().Return(&sql.Tx{}, nil)
	dbService.EXPECT().CloseTx(gomock.Any())

	taskDao := dao.NewMockTaskDao(ctrl)
	taskDao.EXPECT().SaveTask(gomock.Any(), gomock.Any()).Return(nil)

	taskHistoryDao := dao.NewMockTaskHistoryDao(ctrl)
	taskHistoryDao.EXPECT().SaveTaskEntry(gomock.Any(), gomock.Any()).Return(nil)

	taskService := NewTaskService(taskDao, taskHistoryDao, dbService)
	err := taskService.SaveUpdateTask(&task)
	assert.Nil(t, err)
}

func TestTaskService_UpdateTask(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	task := data.Task{
		ID:         "test",
		Name:       "test",
		Status:     data.TASK_UPDATED,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	dbService := NewMockStoreService(ctrl)
	dbService.EXPECT().BeginTx().Return(&sql.Tx{}, nil)
	dbService.EXPECT().CloseTx(gomock.Any())

	taskDao := dao.NewMockTaskDao(ctrl)
	taskDao.EXPECT().SaveTask(gomock.Any(), gomock.Any()).Times(0)

	taskHistoryDao := dao.NewMockTaskHistoryDao(ctrl)
	taskHistoryDao.EXPECT().SaveTaskEntry(gomock.Any(), gomock.Any()).Return(nil)

	taskService := NewTaskService(taskDao, taskHistoryDao, dbService)
	err := taskService.SaveUpdateTask(&task)
	assert.Nil(t, err)
}

func TestTaskService_GetTaskById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"
	taskDTO := model.TaskDTO{
		Id:        id,
		Name:      "test",
		CreatedAt: time.Now(),
	}

	taskHistoryEntry := model.TaskHistoryDTO{
		Id:        1,
		TaskId:    id,
		Status:    data.TASK_CREATED,
		UpdatedAt: taskDTO.CreatedAt,
	}

	dbService := NewMockStoreService(ctrl)
	dbService.EXPECT().BeginTx().Return(&sql.Tx{}, nil)
	dbService.EXPECT().CloseTx(gomock.Any())

	taskDao := dao.NewMockTaskDao(ctrl)
	taskDao.EXPECT().GetTaskById(&id, gomock.Any()).Return(&taskDTO, nil)

	taskHistoryDao := dao.NewMockTaskHistoryDao(ctrl)
	taskHistoryDao.EXPECT().GetTaskById(&id, gomock.Any()).Return(&taskHistoryEntry, nil)

	taskService := NewTaskService(taskDao, taskHistoryDao, dbService)
	res, err := taskService.GetTaskById(&id)

	assert.Nil(t, err)
	assert.Equal(t, res.ID, id)
	assert.Equal(t, res.Name, taskDTO.Name)
	assert.Equal(t, res.Status, taskHistoryEntry.Status)
	assert.Equal(t, res.CreateTime, taskDTO.CreatedAt)
	assert.Equal(t, res.UpdateTime, taskHistoryEntry.UpdatedAt)
}

func TestTaskService_GetTaskByStatus(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	status := data.TASK_UPDATED
	offset := 0
	limit := 1

	taskDTO := model.TaskDTO{
		Id:        "test",
		Name:      "test",
		CreatedAt: time.Now(),
	}

	tasks := make([]*model.TaskDTO, 0)
	tasks = append(tasks, &taskDTO)

	taskHistoryEntry := model.TaskHistoryDTO{
		Id:        1,
		TaskId:    "test",
		Status:    data.TASK_CREATED,
		UpdatedAt: taskDTO.CreatedAt,
	}

	historyEntries := make([]*model.TaskHistoryDTO, 0)
	historyEntries = append(historyEntries, &taskHistoryEntry)

	dbService := NewMockStoreService(ctrl)
	dbService.EXPECT().BeginTx().Return(&sql.Tx{}, nil)
	dbService.EXPECT().CloseTx(gomock.Any())

	taskDao := dao.NewMockTaskDao(ctrl)
	taskDao.EXPECT().GetTasksByIds(gomock.Any(), gomock.Any()).Return(&tasks, nil)

	taskHistoryDao := dao.NewMockTaskHistoryDao(ctrl)
	taskHistoryDao.EXPECT().GetTasksByStatus(&status, &offset, &limit, gomock.Any()).Return(&historyEntries, nil)

	taskService := NewTaskService(taskDao, taskHistoryDao, dbService)
	res, err := taskService.GetTasksByStatus(&status, &offset, &limit)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, res[0].ID, taskDTO.Id)
	assert.Equal(t, res[0].Name, taskDTO.Name)
	assert.Equal(t, res[0].Status, taskHistoryEntry.Status)
	assert.Equal(t, res[0].CreateTime, taskDTO.CreatedAt)
	assert.Equal(t, res[0].UpdateTime, taskHistoryEntry.UpdatedAt)
}

func TestTaskService_GetTaskHistory(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "test"
	taskDTO := model.TaskDTO{
		Id:        id,
		Name:      "test",
		CreatedAt: time.Now(),
	}

	taskHistoryEntry := model.TaskHistoryDTO{
		Id:        1,
		TaskId:    id,
		Status:    data.TASK_CREATED,
		UpdatedAt: taskDTO.CreatedAt,
	}
	historyEntries := make([]*model.TaskHistoryDTO, 0)
	historyEntries = append(historyEntries, &taskHistoryEntry)

	dbService := NewMockStoreService(ctrl)
	dbService.EXPECT().BeginTx().Return(&sql.Tx{}, nil)
	dbService.EXPECT().CloseTx(gomock.Any())

	taskDao := dao.NewMockTaskDao(ctrl)
	taskDao.EXPECT().GetTaskById(&id, gomock.Any()).Return(&taskDTO, nil)

	taskHistoryDao := dao.NewMockTaskHistoryDao(ctrl)
	taskHistoryDao.EXPECT().GetTaskHistory(&id, gomock.Any()).Return(&historyEntries, nil)

	taskService := NewTaskService(taskDao, taskHistoryDao, dbService)
	res, err := taskService.GetTaskHistory(&id)

	assert.Nil(t, err)
	assert.Equal(t, res.ID, id)
	assert.Equal(t, res.Name, taskDTO.Name)
	assert.Equal(t, res.CreateTime, taskDTO.CreatedAt)
	assert.Equal(t, 1, len(res.StatusHistory))
	assert.Equal(t, taskHistoryEntry.Status, res.StatusHistory[0].Status)
	assert.Equal(t, taskHistoryEntry.UpdatedAt, res.StatusHistory[0].UpdatedAt)

}
