package http

import "net/http"

func NewSuccessCreatedResponse() DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusCreated,
		Status:  "SUCCESS",
		Message: "Success Created Data",
	}
}

func NewSuccessDefaultResponse(data interface{}) DefaultResponse {
	return DefaultResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   data,
	}
}
