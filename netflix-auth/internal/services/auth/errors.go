package authentication

import "errors"

var (
	ErrUserWrongPassword = errors.New("wrong password")
	ErrUserNotFound      = errors.New("user with given email not found")
)
