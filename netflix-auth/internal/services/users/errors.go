package users

import "errors"

var ErrUserAlreadyExists = errors.New("user with provided email already exists")
