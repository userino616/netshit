package hash

import (
	"crypto/sha256"
	"fmt"
	"netflix-auth/internal/config"
)

type Service interface {
	IsEqual(userPassword, password string) bool
	Generate(password string) string
}

type hashService struct {
	secret string
}

func NewHashService(cfg *config.Config) Service {
	return &hashService{secret: cfg.JWT.Secret}
}

func (s hashService) IsEqual(hash, str string) bool {
	return hash == s.Generate(str)
}

func (s hashService) Generate(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.secret)))
}
