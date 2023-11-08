package dao

import (
	"database/sql"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/model"
)

type TaskHistoryDao interface {
	SaveTaskEntry(entry *model.TaskHistoryDTO, tx *sql.Tx) error

	GetTaskById(id *string, tx *sql.Tx) (*model.TaskHistoryDTO, error)
	GetTasksByIds(ids *[]string, tx *sql.Tx) (*[]model.TaskHistoryDTO, error)

	GetTasksByStatus(status *data.TaskStatus, tx *sql.Tx) (*[]model.TaskHistoryDTO, error)
	GetTaskHistory(id *string, offset *int, limit *int, tx *sql.Tx) (*[]model.TaskHistoryDTO, error)
}

type TaskHistoryDaoImpl struct {
	db *Database
}

func NewTaskHistoryDao(database *Database) *TaskHistoryDaoImpl {
	return &TaskHistoryDaoImpl{
		db: database,
	}
}

func (dao *TaskHistoryDaoImpl) SaveTaskEntry(entry *model.TaskHistoryDTO, tx *sql.Tx) error {
	return nil
}

// get latest status
func (dao *TaskHistoryDaoImpl) GetTaskById(id *string, tx *sql.Tx) (*model.TaskHistoryDTO, error) {
	return nil, nil
}

// get latest status
func (dao *TaskHistoryDaoImpl) GetTasksByIds(ids *[]string, tx *sql.Tx) (*[]model.TaskHistoryDTO, error) {
	return nil, nil
}

func (dao *TaskHistoryDaoImpl) GetTasksByStatus(status *data.TaskStatus, tx *sql.Tx) (*[]model.TaskHistoryDTO, error) {
	return nil, nil
}

func (dao *TaskHistoryDaoImpl) GetTaskHistory(id *string, offset *int, limit *int, tx *sql.Tx) (*[]model.TaskHistoryDTO, error) {
	return nil, nil
}
