package httperror

import (
	"errors"
	"net/http"
	"netflix-auth/pkg/logger"
)

const (
	InternalServerError = "Internal Server Error"
)

var ErrType = errors.New("error is not of type HTTPError")

type HTTPError interface {
	Error() string
	Unwrap() error
	GetMessage() string
}

type Wrapper struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
	Err     error  `json:"-"`
}

func NewHTTPErrorWrapper(code int, err error, msg string) HTTPError {
	if msg == "" {
		msg = "Unknown error"
	}

	return Wrapper{
		Message: msg,
		Code:    code,
		Err:     err,
	}
}

func (err Wrapper) Error() string {
	if err.Err != nil {
		return err.Err.Error()
	}

	return err.Message
}

func (err Wrapper) Unwrap() error {
	return err.Err
}

func NewBadRequestErr(err error, msg string) HTTPError {
	return NewHTTPErrorWrapper(http.StatusBadRequest, err, msg)
}

func NewInternalServerErr(err error) HTTPError {
	l := logger.GetLogger()
	l.Error(err)

	return NewHTTPErrorWrapper(http.StatusInternalServerError, err, InternalServerError)
}

func NewForbiddenErr(err error, msg string) HTTPError {
	return NewHTTPErrorWrapper(http.StatusForbidden, err, msg)
}

func NewNotFoundErr(err error, msg string) HTTPError {
	return NewHTTPErrorWrapper(http.StatusNotFound, err, msg)
}

func (err Wrapper) GetMessage() string {
	return err.Message
}
