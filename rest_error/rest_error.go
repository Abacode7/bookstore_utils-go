package rest_error

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{Message: message,
		Status: http.StatusBadRequest,
		Error:  http.StatusText(http.StatusBadRequest),
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   http.StatusText(http.StatusNotFound),
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   http.StatusText(http.StatusInternalServerError),
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusUnauthorized,
		Error: http.StatusText(http.StatusUnauthorized),
	}
}
