package helper

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrDataAlreadyExist    = errors.New("data already exist")
	ErrStatusInternalError = errors.New("internal server error")
	ErrUserNotFound        = errors.New("user not found")
	ErrWrongPassword       = errors.New("wrong password")
)
