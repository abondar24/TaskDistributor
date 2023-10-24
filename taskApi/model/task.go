package model

type Task struct {
	ID         string
	Name       string
	Status     TaskStatus
	CreateTime string
	UpdateTime string
}

type TaskStatus = string

const (
	TASK_CREATED TaskStatus = "created"
	TASK_UPDATED TaskStatus = "updated"
	TASK_DELETED TaskStatus = "deleted"
)
