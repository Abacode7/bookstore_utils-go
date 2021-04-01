package rest_error

import (
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
}

type restErr struct {
	message string `json:"message"`
	status  int    `json:"status"`
	error   string `json:"error"`
}

func (re *restErr) Message() string {
	return re.message
}

func (re *restErr) Status() int {
	return re.status
}

func (re *restErr) Error() string {
	return fmt.Sprintf(
		"message: %s; status: %d; error; %s",
		re.message,
		re.status,
		re.error,
	)
}

func NewBadRequestError(message string) RestErr {
	return &restErr{
		message: message,
		status: http.StatusBadRequest,
		error:  http.StatusText(http.StatusBadRequest),
	}
}

func NewNotFoundError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusNotFound,
		error:   http.StatusText(http.StatusNotFound),
	}
}

func NewInternalServerError(message string) RestErr {
	return &restErr{
		message: message,
		status:  http.StatusInternalServerError,
		error:   http.StatusText(http.StatusInternalServerError),
	}
}

func NewUnauthorizedError(message string) RestErr {
	return &restErr{
		message: message,
		status: http.StatusUnauthorized,
		error: http.StatusText(http.StatusUnauthorized),
	}
}
