package jwt

import (
	"netflix-auth/internal/config"
	"netflix-auth/internal/repository"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	httperror "netflix-auth/pkg/http_error"
)

type Service interface {
	GenerateAccessToken(userID uuid.UUID) (string, httperror.HTTPError)
	ParseToken(accessToken string) (*TokenClaims, httperror.HTTPError)

	AddToBlackList(tokenID string, exp time.Duration) httperror.HTTPError
	IsBlacklisted(tokenID string) (bool, httperror.HTTPError)
}

type service struct {
	tokenExp time.Duration
	secret   string
	inMemory repository.MemStore
}

func NewJWTService(cfg *config.Config, memStore repository.MemStore) Service {
	return &service{
		tokenExp: time.Duration(cfg.JWT.AccessTokenExpiryHours) * time.Hour,
		secret:   cfg.JWT.Secret,
		inMemory: memStore,
	}
}

func (s service) GenerateAccessToken(userID uuid.UUID) (string, httperror.HTTPError) {
	claims := newTokenClaims(userID, s.tokenExp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return ss, httperror.NewInternalServerErr(err)
	}

	return ss, nil
}

func (s service) ParseToken(accessToken string) (*TokenClaims, httperror.HTTPError) {
	tokenParts := strings.Split(accessToken, " ")
	if len(tokenParts) != 2 {
		return nil, httperror.NewBadRequestErr(nil, "wrong access token")
	}
	token, err := jwt.ParseWithClaims(tokenParts[1], &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httperror.NewInternalServerErr(ErrInvalidSingMethod)
		}

		return []byte(s.secret), nil
	})
	if err != nil {
		if isExpired(err) {
			return nil, httperror.NewForbiddenErr(err, err.Error())
		}

		return nil, httperror.NewInternalServerErr(err)
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, httperror.NewInternalServerErr(ErrTokenWrongType)
	}
	return claims, nil
}

func (s service) AddToBlackList(tokenID string, exp time.Duration) httperror.HTTPError {
	err := s.inMemory.BlackListJWT(tokenID, exp)
	if err != nil {
		return httperror.NewInternalServerErr(err)
	}
	return nil
}

func (s service) IsBlacklisted(tokenID string) (bool, httperror.HTTPError) {
	val, err := s.inMemory.IsJWTBlacklisted(tokenID)
	if err != nil {
		return val, httperror.NewInternalServerErr(err)
	}
	return val, nil
}
