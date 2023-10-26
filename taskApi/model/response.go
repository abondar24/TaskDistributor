package model

type SendResponse struct {
	ID string `json:"taskId"`
}

type ErrorResponse struct {
	ERROR string `json:"error"`
}

type HealthResponse struct {
	MESSAGE string `json:"message"`
}
