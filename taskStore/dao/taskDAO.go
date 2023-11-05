package dao

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/model"
)

type TaskDao interface {
	SaveTask(task *data.Task) (*model.TaskDTO, error)
	GetTaskById(id *string) (*model.TaskDTO, error)
	GetTasks(offset *int, limit *int) (*[]model.TaskDTO, error)
}

type TaskDaoImpl struct {
	db *Database
}

func NewTaskDao(database *Database) *TaskDaoImpl {
	return &TaskDaoImpl{
		db: database,
	}
}

func (dao *TaskDaoImpl) SaveTask(task *data.Task) (*model.TaskDTO, error) {
	return nil, nil
}

func (dao *TaskDaoImpl) GetTaskById(id *string) (*model.TaskDTO, error) {
	return nil, nil
}
func (dao *TaskDaoImpl) GetTasks(offset *int, limit *int) (*[]model.TaskDTO, error) {
	return nil, nil
}
