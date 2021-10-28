package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWT struct {
	Token string `json:"accessToken"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"userId"`
}

func newTokenClaims(userID uuid.UUID, tokenExp time.Duration) *TokenClaims {
	return &TokenClaims{
		jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: time.Now().Add(tokenExp).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	}
}
