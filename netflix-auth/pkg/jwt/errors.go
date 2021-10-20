package jwt

import "errors"

var (
	ErrInvalidSingMethod = errors.New("invalid signing method")
	ErrTokenWrongType    = errors.New("token claims are not of type *tokenClaims")
)
