package models

import (
	"net/http"
)

type RestErr struct {
	Message string
	Status  int
	Error   string
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Interal Server Error",
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
