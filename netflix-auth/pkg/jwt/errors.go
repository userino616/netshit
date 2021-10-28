package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidSingMethod = errors.New("invalid signing method")
	ErrTokenWrongType    = errors.New("token claims are not of type *TokenClaims")
)

func isExpired(err error) bool {
	jwtError, ok := err.(*jwt.ValidationError)
	if ok {
		return jwtError.Errors == jwt.ValidationErrorExpired
	}
	return false
}
