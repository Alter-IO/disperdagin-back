package http

import "net/http"

func NewSuccessCreatedResponse(message string) DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusCreated,
		Status:  "SUCCESS",
		Message: message,
	}
}

func NewCreatedResponseWithData(data interface{}) DefaultResponse {
	return DefaultResponse{
		Code:   http.StatusCreated,
		Status: "SUCCESS",
		Data:   data,
	}
}

func NewSuccessDefaultResponse(data interface{}, message string) DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusOK,
		Status:  "SUCCESS",
		Message: message,
		Data:    data,
	}
}
