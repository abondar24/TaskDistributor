package model

type TaskResponse struct {
	ID string `json:"taskId"`
}

type ErrorResponse struct {
	ERROR string `json:"error"`
}

type HealthResponse struct {
	MESSAGE string `json:"message"`
}
