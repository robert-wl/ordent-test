package utils

type ErrorResponse struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Error      []string `json:"error"`
}

func NewErrorResponse(message string, code int, errors ...string) *ErrorResponse {
	return &ErrorResponse{
		code,
		message,
		errors,
	}
}

type AppError struct {
	Type    string
	Message string
	Err     error
}
