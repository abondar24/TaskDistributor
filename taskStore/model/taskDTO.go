package model

import "github.com/abondar24/TaskDistributor/taskData/data"

type TaskDTO struct {
	TaskId    string
	Name      string
	CreatedAt string
}

type TaskHistoryDTO struct {
	ID        string
	Status    data.TaskStatus
	UpdatedAt string
}
