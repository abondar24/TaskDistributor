package dao

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/model"
)

type TaskHistoryDao interface {
	SaveTaskEntry(entry *model.TaskHistoryDTO) error
	GetTasksByStatus(status *data.TaskStatus) (*[]model.TaskHistoryDTO, error)
	GetTaskHistory(id *string, offset *int, limit *int) (*[]model.TaskHistoryDTO, error)
}

type TaskHistoryDaoImpl struct {
	db *Database
}

func NewTaskHistoryDao(database *Database) *TaskHistoryDaoImpl {
	return &TaskHistoryDaoImpl{
		db: database,
	}
}

func (dao *TaskHistoryDaoImpl) SaveTaskEntry(entry *model.TaskHistoryDTO) error {
	return nil
}

func (dao *TaskHistoryDaoImpl) GetTasksByStatus(status *data.TaskStatus) (*[]model.TaskHistoryDTO, error) {
	return nil, nil
}

func (dao *TaskHistoryDaoImpl) GetTaskHistory(id *string, offset *int, limit *int) (*[]model.TaskHistoryDTO, error) {
	return nil, nil
}
