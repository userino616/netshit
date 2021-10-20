package postgres

import "github.com/go-pg/pg/v10"

const (
	duplicateErrorCode = "23505" // не нашел я такой ошибки в пакете pg
)

func DuplicateError(err error) bool {
	if err == nil {
		return false
	}

	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		return pgErr.Field('C') == duplicateErrorCode
	}

	return false
}
