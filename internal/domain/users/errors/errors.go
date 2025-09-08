package errors

import "errors"

var (
	ErrorUserIsAlreadyExist = errors.New("user with same email or phone is already exist")
	ErrorInvalidPassword    = errors.New("wrong password")
	ErrorUserNotExist       = errors.New("user not found")
)
