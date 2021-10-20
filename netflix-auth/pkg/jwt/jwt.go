package jwt

import (
	"netflix-auth/internal/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	httperror "netflix-auth/pkg/http_error"
)

type Service interface {
	GenerateAccessToken(userId uuid.UUID) (string, httperror.HTTPError)
	ParseToken(accessToken string) (uuid.UUID, httperror.HTTPError)
}

type service struct {
	tokenExp time.Duration
	secret   string
}

func NewJWTService(cfg *config.Config) Service {
	return &service{
		tokenExp: time.Duration(cfg.JWT.AccessTokenExpiryHours) * time.Hour,
		secret:   cfg.JWT.Secret,
	}
}

func (s service) GenerateAccessToken(userId uuid.UUID) (string, httperror.HTTPError) {
	claims := newTokenClaims(userId, s.tokenExp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return ss, httperror.NewInternalServerErr(err)
	}

	return ss, nil
}

func (s service) ParseToken(accessToken string) (uuid.UUID, httperror.HTTPError) {
	tokenParts := strings.Split(accessToken, " ")
	token, err := jwt.ParseWithClaims(tokenParts[1], &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httperror.NewInternalServerErr(ErrInvalidSingMethod)
		}

		return []byte(s.secret), nil
	})
	if err != nil {
		return uuid.UUID{}, httperror.NewInternalServerErr(err)
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return uuid.UUID{}, httperror.NewInternalServerErr(ErrTokenWrongType)
	}

	return claims.UserID, nil
}
