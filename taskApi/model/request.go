package model

// TaskRequest Request to create a new task
type TaskRequest struct {
	Name string `json:"taskName"`
}
