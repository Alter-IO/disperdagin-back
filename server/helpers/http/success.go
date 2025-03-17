package http

import "net/http"

func NewSuccessCreatedResponse(message string) DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusCreated,
		Status:  "SUCCESS",
		Message: message,
	}
}

func NewSuccessDefaultResponse(data interface{}, message string) DefaultResponse {
	return DefaultResponse{
		Code:   http.StatusOK,
		Status: message,
		Data:   data,
	}
}
