package dao

import (
	"database/sql"
	"fmt"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"log"
)

type TaskDao interface {
	SaveTask(task *model.TaskDTO) error

	GetTaskById(id *string) (*model.TaskDTO, error)

	GetTasksByIds(ids []*string) (*[]*model.TaskDTO, error)
}

type TaskDaoImpl struct {
	db *Database
}

func NewTaskDao(database *Database) *TaskDaoImpl {
	return &TaskDaoImpl{
		db: database,
	}
}

// TODO we need to run transactions in service not here
func (dao *TaskDaoImpl) SaveTask(task *model.TaskDTO) error {
	tx, err := dao.db.BeginTx()
	if err != nil {
		return err
	}

	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Println(err.Error())
		}
	}(tx)

	query := fmt.Sprintf("INSERT INTO task(id,name,created_at) VALUES ('%v','%v','%v')", task.Id, task.Name, task.CreatedAt)
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
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

func (dao *TaskDaoImpl) GetTaskById(id *string) (*model.TaskDTO, error) {
	tx, err := dao.db.BeginTx()
	if err != nil {
		return nil, err
	}

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
func (dao *TaskDaoImpl) GetTasksByIds(ids []*string) (*[]*model.TaskDTO, error) {
	tx, err := dao.db.BeginTx()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM task WHERE id IN (%v)", ids)

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
		return nil, err
	}

	for rows.Next() {
		task := model.TaskDTO{}
		err := rows.Scan(&task.Id, &task.Name, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &task)
	}

	return &result, nil
}
