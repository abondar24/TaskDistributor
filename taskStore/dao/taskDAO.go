package dao

import (
	"database/sql"
	"fmt"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"log"
)

type TaskDao interface {
	SaveTask(task *model.TaskDTO, tx *sql.Tx) error

	GetTaskById(id *string, tx *sql.Tx) (*model.TaskDTO, error)

	GetTasksByIds(ids []*string, tx *sql.Tx) (*[]*model.TaskDTO, error)
}

type TaskDaoImpl struct {
}

func NewTaskDao() *TaskDaoImpl {
	return &TaskDaoImpl{}
}

func (dao *TaskDaoImpl) SaveTask(task *model.TaskDTO, tx *sql.Tx) error {

	query := fmt.Sprintf("INSERT INTO task(id,name,created_at) VALUES ('%v','%v','%v')", task.Id, task.Name, task.CreatedAt.Format(model.TimeFormat))
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	return nil
}

func (dao *TaskDaoImpl) GetTaskById(id *string, tx *sql.Tx) (*model.TaskDTO, error) {

	query := fmt.Sprintf("SELECT * FROM task WHERE id='%v'", *id)

	result := &model.TaskDTO{}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var createdAtRaw sql.RawBytes
		err := rows.Scan(&result.Id, &result.Name, &createdAtRaw)
		if err != nil {
			return nil, err
		}

		result.CreatedAt, err = ConvertTime(createdAtRaw)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
func (dao *TaskDaoImpl) GetTasksByIds(ids []*string, tx *sql.Tx) (*[]*model.TaskDTO, error) {

	query := fmt.Sprintf("SELECT * FROM task WHERE id IN (%v)", ids)

	result := make([]*model.TaskDTO, 0)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := model.TaskDTO{}
		var createdAtRaw sql.RawBytes

		err := rows.Scan(&task.Id, &task.Name, &createdAtRaw)
		if err != nil {
			return nil, err
		}

		task.CreatedAt, err = ConvertTime(createdAtRaw)
		if err != nil {
			return nil, err
		}

		result = append(result, &task)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	return &result, nil
}
