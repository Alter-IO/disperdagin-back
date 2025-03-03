package http

import (
	"alter-io-go/helpers/derrors"
	"alter-io-go/helpers/logger"
	"errors"
	"log/slog"
	"net/http"
)

const (
	ErrBadRequest     = "BAD_REQUEST"
	ErrValidation     = "VALIDATION_ERROR"
	ErrInternalServer = "SERVER_ERROR"
	ErrNotFound       = "NOT_FOUND"
	ErrUnauthorized   = "UNAUTHORIZED"
	ErrForbidden      = "FORBIDDEN"
	ErrConflict       = "CONFLICT"
)

type DefaultResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code     int    `json:"code"`
	Status   string `json:"status,omitempty"`
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

func NewBadRequestResponse(errMsg string) DefaultResponse {
	return DefaultResponse{
		Code:    400,
		Status:  ErrBadRequest,
		Message: errMsg,
	}
}

func NewValidationResponse(errMsg string) DefaultResponse {
	return DefaultResponse{
		Code:    400,
		Status:  ErrValidation,
		Message: errMsg,
	}
}

func NewUnauthorizedResponse(msg string) DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusUnauthorized,
		Status:  ErrUnauthorized,
		Message: msg,
	}
}

func NewForbiddenResponse() DefaultResponse {
	return DefaultResponse{
		Code:    http.StatusForbidden,
		Status:  ErrForbidden,
		Message: "Forbidden",
	}
}

func MapErrorToResponse(err error) ErrorResponse {
	defaultErr := ErrorResponse{Code: http.StatusInternalServerError, Status: ErrInternalServer, Message: "Internal Server Error"}
	var ierr *derrors.Error
	if !errors.As(err, &ierr) {
		logger.Get().Debug("Internal Error", slog.String("error", err.Error()))
		return defaultErr
	} else {
		logger.Get().Debug("Internal Error", slog.String("error", ierr.Error()))
		cases := map[derrors.ErrorCode]ErrorResponse{
			derrors.ErrorCodeBadRequest:   {Code: http.StatusBadRequest, Status: ErrBadRequest, Message: ierr.Error(), Internal: ierr},
			derrors.ErrorCodeNotFound:     {Code: http.StatusNotFound, Status: ErrNotFound, Message: ierr.Error(), Internal: ierr},
			derrors.ErrorCodeUnauthorized: {Code: http.StatusUnauthorized, Status: ErrUnauthorized, Message: "Unauthorized", Internal: ierr},
			derrors.ErrorCodeDuplicate:    {Code: http.StatusConflict, Status: ErrConflict, Message: ierr.Error(), Internal: ierr},
			derrors.ErrorCodeForbidden:    {Code: http.StatusForbidden, Status: ErrForbidden, Message: ierr.Error(), Internal: ierr},
		}

		catchErr, found := cases[ierr.Code()]
		if found {
			return catchErr
		}
	}
	return defaultErr
}
