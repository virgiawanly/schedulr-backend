package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppError struct {
	Code    codes.Code
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) ToGRPCError() error {
	return status.Error(e.Code, e.Message)
}

func NewAppError(code codes.Code, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func WrapError(code codes.Code, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: err.Error(),
	}
}
