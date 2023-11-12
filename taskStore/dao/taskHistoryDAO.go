package dao

import (
	"database/sql"
	"fmt"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"log"
)

type TaskHistoryDao interface {
	SaveTaskEntry(entry *model.TaskHistoryDTO) error

	GetTaskById(id *string) (*model.TaskHistoryDTO, error)
	GetTasksByStatus(status *data.TaskStatus, offset *int, limit *int) (*[]*model.TaskHistoryDTO, error)
	GetTaskHistory(id *string) (*[]*model.TaskHistoryDTO, error)
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
	tx, err := dao.db.BeginTx()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO task_history(task_id,status,updated_at)  VALUES ('%v','%v','%v')", entry.TaskId, entry.Status, entry.UpdatedAt)
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

// get latest status
func (dao *TaskHistoryDaoImpl) GetTaskById(id *string) (*model.TaskHistoryDTO, error) {
	tx, err := dao.db.BeginTx()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM task_history WHERE task_id='%v' ORDER BY updated_at ", id)

	result := &model.TaskHistoryDTO{}

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
		err := rows.Scan(&result.Id, &result.TaskId, &result.Status, &result.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (dao *TaskHistoryDaoImpl) GetTasksByStatus(status *data.TaskStatus,
	offset *int, limit *int) (*[]*model.TaskHistoryDTO, error) {
	tx, err := dao.db.BeginTx()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM task_history WHERE status='%v' LIMIT %v OFFSET %v", status, limit, offset)

	return fetchTasks(query, tx)
}

func (dao *TaskHistoryDaoImpl) GetTaskHistory(id *string) (*[]*model.TaskHistoryDTO, error) {
	tx, err := dao.db.BeginTx()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("SELECT * FROM task_history WHERE id='%v'", id)

	return fetchTasks(query, tx)
}

func fetchTasks(query string, tx *sql.Tx) (*[]*model.TaskHistoryDTO, error) {
	result := make([]*model.TaskHistoryDTO, 0)

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
		entry := model.TaskHistoryDTO{}
		err := rows.Scan(&entry.Id, &entry.TaskId, &entry.Status, &entry.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &entry)
	}

	return &result, nil
}
