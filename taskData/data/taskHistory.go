package data

import "time"

type TaskHistory struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	CreateTime    time.Time         `json:"createdAt"`
	StatusHistory []TaskStatusEntry `json:"statusHistory"`
}

type TaskStatusEntry struct {
	Status    TaskStatus `json:"status"`
	UpdatedAt time.Time  `json:"updated_at"`
}
