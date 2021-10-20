package passwords

import (
	"crypto/sha256"
	"fmt"
	"netflix-auth/internal/config"

	httperror "netflix-auth/pkg/http_error"
)

type Service interface {
	CompareHash(userPassword, password string) httperror.HTTPError
	GeneratePasswordHash(password string) string
}

type passwordService struct {
	secret string
}

func NewPasswordService(cfg *config.Config) Service {
	return &passwordService{secret: cfg.JWT.Secret}
}

func (s passwordService) CompareHash(userPassword, password string) httperror.HTTPError {
	if userPassword != s.GeneratePasswordHash(password) {
		return httperror.NewForbiddenErr(nil, "wrong password")
	}

	return nil
}

func (s passwordService) GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.secret)))
}
