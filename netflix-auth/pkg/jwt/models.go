package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWT struct {
	Token string `json:"accessToken"`
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"userId"`
}

func newTokenClaims(userId uuid.UUID, tokenExp time.Duration) *tokenClaims {
	return &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExp).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	}
}
