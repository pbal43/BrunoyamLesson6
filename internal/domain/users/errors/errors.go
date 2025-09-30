package errors

import "errors"

var (
	ErrUserIsAlreadyExist = errors.New("user with same email or phone is already exist")
	ErrInvalidPassword    = errors.New("wrong password")
	ErrUserNotExist       = errors.New("user not found")
)
