package errors

import "errors"

var (
	ErrDBConnection   = errors.New("failed to connect to database")
	ErrDBQuery        = errors.New("database query error")
	ErrDBTransaction  = errors.New("database transaction error")
	ErrRecordExists   = errors.New("record already exists")
	ErrRecordNotFound = errors.New("record not found")
)
