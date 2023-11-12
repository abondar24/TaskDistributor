package data

import "time"

type Task struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Status     TaskStatus `json:"status"`
	CreateTime time.Time  `json:"createdAt"`
	UpdateTime time.Time  `json:"updatedAt"`
}

type TaskStatus = string

const (
	TASK_CREATED TaskStatus = "created"
	TASK_UPDATED TaskStatus = "updated"
	TASK_DELETED TaskStatus = "deleted"

	TASK_COMPLETED TaskStatus = "completed"
)
