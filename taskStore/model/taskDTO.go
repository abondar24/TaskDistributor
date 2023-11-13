package model

import (
	"github.com/abondar24/TaskDistributor/taskData/data"
	"time"
)

type TaskDTO struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type TaskHistoryDTO struct {
	Id        int64           `db:"id"`
	TaskId    string          `db:"task_id"`
	Status    data.TaskStatus `db:"status"`
	UpdatedAt time.Time       `db:"updated_at"`
}

const (
	TimeFormat = "2006-01-02 15:04:05"
)
