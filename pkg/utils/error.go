package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

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
	Err     error
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(err error, code int, message string) *AppError {
	return &AppError{
		Err:     err,
		Code:    code,
		Message: message,
	}
}

func SendError(ctx *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		ctx.JSON(
			appErr.Code,
			NewErrorResponse(
				http.StatusText(appErr.Code),
				appErr.Code,
				appErr.Message,
			),
		)
		return
	}

	if errors.As(err, &validator.ValidationErrors{}) {
		ctx.JSON(
			http.StatusBadRequest,
			NewErrorResponse(
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
				err.Error(),
			),
		)
		return
	}

	log.Println("An unexpected error occurred: ", err)
	ctx.JSON(
		http.StatusInternalServerError,
		NewErrorResponse(
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
			"An unexpected error occurred. Please try again later.",
		),
	)
}
