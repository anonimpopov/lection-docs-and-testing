package errors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrTokenInvalid = errors.New("invalid token")
)