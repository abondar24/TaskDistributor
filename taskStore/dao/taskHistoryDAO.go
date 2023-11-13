package dao

import (
	"database/sql"
	"fmt"
	"github.com/abondar24/TaskDistributor/taskData/data"
	"github.com/abondar24/TaskDistributor/taskStore/model"
	"log"
)

type TaskHistoryDao interface {
	SaveTaskEntry(entry *model.TaskHistoryDTO, tx *sql.Tx) error

	GetTaskById(id *string, tx *sql.Tx) (*model.TaskHistoryDTO, error)
	GetTasksByStatus(status *data.TaskStatus, offset *int, limit *int, tx *sql.Tx) (*[]*model.TaskHistoryDTO, error)
	GetTaskHistory(id *string, tx *sql.Tx) (*[]*model.TaskHistoryDTO, error)
}

type TaskHistoryDaoImpl struct {
}

func NewTaskHistoryDao() *TaskHistoryDaoImpl {
	return &TaskHistoryDaoImpl{}
}

func (dao *TaskHistoryDaoImpl) SaveTaskEntry(entry *model.TaskHistoryDTO, tx *sql.Tx) error {

	query := fmt.Sprintf("INSERT INTO task_history(task_id,status,updated_at)  VALUES ('%v','%v','%v')", entry.TaskId, entry.Status, entry.UpdatedAt.Format(model.TimeFormat))
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

	return nil
}

// get latest status
func (dao *TaskHistoryDaoImpl) GetTaskById(id *string, tx *sql.Tx) (*model.TaskHistoryDTO, error) {

	query := fmt.Sprintf("SELECT * FROM task_history WHERE task_id='%v' ORDER BY updated_at ", *id)

	result := &model.TaskHistoryDTO{}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var updatedAtRaw sql.RawBytes
		err := rows.Scan(&result.Id, &result.TaskId, &result.Status, &updatedAtRaw)
		if err != nil {
			return nil, err
		}

		result.UpdatedAt, err = ConvertTime(updatedAtRaw)
		if err != nil {
			return nil, err
		}
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	return result, nil
}

func (dao *TaskHistoryDaoImpl) GetTasksByStatus(status *data.TaskStatus,
	offset *int, limit *int, tx *sql.Tx) (*[]*model.TaskHistoryDTO, error) {
	query := fmt.Sprintf("SELECT * FROM task_history WHERE status='%v' LIMIT %v OFFSET %v", *status, *limit, *offset)

	return fetchTasks(query, tx)
}

func (dao *TaskHistoryDaoImpl) GetTaskHistory(id *string, tx *sql.Tx) (*[]*model.TaskHistoryDTO, error) {
	query := fmt.Sprintf("SELECT * FROM task_history WHERE id='%v'", id)

	return fetchTasks(query, tx)
}

func fetchTasks(query string, tx *sql.Tx) (*[]*model.TaskHistoryDTO, error) {
	result := make([]*model.TaskHistoryDTO, 0)

	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		entry := model.TaskHistoryDTO{}
		var updatedAtRaw sql.RawBytes

		err := rows.Scan(&entry.Id, &entry.TaskId, &entry.Status, &updatedAtRaw)
		if err != nil {
			return nil, err
		}

		entry.UpdatedAt, err = ConvertTime(updatedAtRaw)
		if err != nil {
			return nil, err
		}

		result = append(result, &entry)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	return &result, nil
}
