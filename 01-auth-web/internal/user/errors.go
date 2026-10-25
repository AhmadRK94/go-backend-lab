package user

import "errors"

var (
	ErrDuplicateEmail = errors.New("email already exist")
	ErrUserNotFound   = errors.New("user not found")
)
