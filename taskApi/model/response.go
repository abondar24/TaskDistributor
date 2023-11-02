package model

// TaskResponse Id of created task
type TaskResponse struct {
	ID string `json:"taskId"`
}

// ErrorResponse response in case of error
type ErrorResponse struct {
	ERROR string `json:"error"`
}
