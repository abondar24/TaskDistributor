package data

type TaskHistory struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	CreateTime    string            `json:"createdAt"`
	StatusHistory []TaskStatusEntry `json:"statusHistory"`
}

type TaskStatusEntry struct {
	Status    TaskStatus `json:"status"`
	UpdatedAt string     `json:"updated_at"`
}
