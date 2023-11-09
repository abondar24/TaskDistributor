package model

import "github.com/abondar24/TaskDistributor/taskData/data"

type TaskDTO struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"created_at"`
}

type TaskHistoryDTO struct {
	Id        int64           `db:"id"`
	TaskId    string          `db:"task_id"`
	Status    data.TaskStatus `db:"status"`
	UpdatedAt string          `db:"updated_at"`
}
