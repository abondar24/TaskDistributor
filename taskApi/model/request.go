package model

type TaskRequest struct {
	Name string `json:"taskName"`
	ID   string
}
