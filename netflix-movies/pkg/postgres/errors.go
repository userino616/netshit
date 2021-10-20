package postgres

import (
	"github.com/go-pg/pg/v10"
)

const (
	DuplicateErrorCode = "23505" // не нашел я такой ошибки в пакете pg
	ViolatesForeignKeyCode = "23503"
)

// var ErrAlreadyExists = errors.New("record already exists")

func GetPgErrCode(err error) string {
	if err == nil {
		return ""
	}

	pgErr, ok := err.(pg.Error)
	if ok {
		return pgErr.Field('C')
	}

	return ""
}

func IsDuplicateError(err error) bool {
	return GetPgErrCode(err) == DuplicateErrorCode
}

func IsViolatesForeignKeyError(err error) bool {
	return GetPgErrCode(err) == ViolatesForeignKeyCode
}
