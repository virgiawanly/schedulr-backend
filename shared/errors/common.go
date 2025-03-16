package errors

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("access forbidden")
	ErrConflict       = errors.New("resource already exists")
	ErrInternalServer = errors.New("internal server error")
)
