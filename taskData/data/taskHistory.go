package data

import "time"

// TaskHistory task with history of status changes
//
// swagger:model
type TaskHistory struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	CreateTime    time.Time         `json:"createdAt"`
	StatusHistory []TaskStatusEntry `json:"statusHistory"`
}

// TaskStatusEntry task status change entry
//
// swagger:model
type TaskStatusEntry struct {
	Status    TaskStatus `json:"status"`
	UpdatedAt time.Time  `json:"updated_at"`
}
