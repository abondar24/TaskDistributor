package response

// HealthResponse response to healthcheck
type HealthResponse struct {
	MESSAGE string `json:"message"`
}

// ErrorResponse response in case of error
type ErrorResponse struct {
	ERROR string `json:"error"`
}
