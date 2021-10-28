package auth

import (
	"netflix-auth/internal/models"
	"netflix-auth/internal/services/users"
	"netflix-auth/pkg/hash"
	httperror "netflix-auth/pkg/http_error"
	"netflix-auth/pkg/jwt"
	"time"
)

type contextKey string

const (
	UserIDKey      contextKey = "userID"
	TokenClaimsKey contextKey = "tokenClaims"
)

type Service interface {
	Authenticate(data models.User) (*jwt.JWT, httperror.HTTPError)
	LogOut(token string, exp int64) httperror.HTTPError
}

type authService struct {
	userService users.Service
	jwtService  jwt.Service
	hashService hash.Service
}

func NewAuthService(us users.Service, js jwt.Service, hs hash.Service) Service {
	return &authService{
		userService: us,
		jwtService:  js,
		hashService: hs,
	}
}

func (s authService) Authenticate(data models.User) (*jwt.JWT, httperror.HTTPError) {
	user, err := s.userService.GetByEmail(data.Email)
	if err != nil {
		return nil, httperror.NewBadRequestErr(ErrUserNotFound, ErrUserNotFound.Error())
	}

	passwordIsEqual := s.hashService.IsEqual(user.Password, data.Password)
	if !passwordIsEqual {
		return nil, httperror.NewBadRequestErr(ErrUserWrongPassword, ErrUserWrongPassword.Error())
	}

	token, httpError := s.jwtService.GenerateAccessToken(user.ID)
	if httpError != nil {
		return nil, httpError
	}

	return &jwt.JWT{Token: token}, nil
}

func (s authService) LogOut(token string, exp int64) httperror.HTTPError {
	expTime := time.Unix(exp, 0)
	timeLeft := expTime.Sub(time.Now())
	return s.jwtService.AddToBlackList(token, timeLeft)
}
