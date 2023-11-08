package dao

import (
	"database/sql"
	"fmt"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"log"
)

type TaskDao interface {
	SaveTask(task *data.Task, tx *sql.Tx) error

	GetTaskById(id *string, tx *sql.Tx) (*model.TaskDTO, error)
	GetTasks(offset *int, limit *int, tx *sql.Tx) (*[]*model.TaskDTO, error)
}

type TaskDaoImpl struct {
	db *Database
}

func NewTaskDao(database *Database) *TaskDaoImpl {
	return &TaskDaoImpl{
		db: database,
	}
}

func (dao *TaskDaoImpl) SaveTask(task *data.Task, tx *sql.Tx) error {

	query := fmt.Sprintf("INSERT INTO task(id,name,created_at) VALUES ('%v','%v','%v')", task.ID, task.Name, task.CreateTime)
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (dao *TaskDaoImpl) GetTaskById(id *string, tx *sql.Tx) (*model.TaskDTO, error) {

	query := fmt.Sprintf("SELECT * FROM task WHERE id='%v'", id)

	result := &model.TaskDTO{}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&result.Id, &result.Name, &result.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
func (dao *TaskDaoImpl) GetTasks(offset *int, limit *int, tx *sql.Tx) (*[]*model.TaskDTO, error) {

	query := fmt.Sprintf("SELECT * FROM task LIMIT %v OFFSET %v", limit, offset)

	result := make([]*model.TaskDTO, 0)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		task := model.TaskDTO{}
		err := rows.Scan(&task.Id, &task.Name, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}
