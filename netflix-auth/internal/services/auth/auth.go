package authentication

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"net/http"
	"netflix-auth/internal/models"
	"netflix-auth/internal/repository/users"
	"netflix-auth/pkg/jwt"
	"netflix-auth/pkg/passwords"

	"github.com/google/uuid"

	httperror "netflix-auth/pkg/http_error"
)

type AuthService interface {
	Authenticate(data models.User) (*jwt.JWT, httperror.HTTPError)
	Middleware(next http.Handler) http.Handler
	GetUserByEmail(email string) (*models.User, httperror.HTTPError)
	GetUserByID(id uuid.UUID) (*models.User, httperror.HTTPError)
}

type authService struct {
	userRepo        users.UserRepository
	jwtService      jwt.Service
	passwordService passwords.Service
}

func NewAuthService(r users.UserRepository, js jwt.Service, ps passwords.Service) AuthService {
	return &authService{
		userRepo:        r,
		jwtService:      js,
		passwordService: ps,
	}
}

func (s authService) GetUserByID(id uuid.UUID) (*models.User, httperror.HTTPError) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, httperror.NewNotFoundErr(err, "user not found")
		}
		return nil, httperror.NewInternalServerErr(err)
	}

	return &user, nil
}

func (s authService) GetUserByEmail(email string) (*models.User, httperror.HTTPError) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, httperror.NewNotFoundErr(err, "user not found")
		}
		return nil, httperror.NewInternalServerErr(err)
	}

	return &user, nil
}

func (s authService) Authenticate(data models.User) (*jwt.JWT, httperror.HTTPError) {
	user, err := s.GetUserByEmail(data.Email)
	if err != nil {
		// TODO check if error is NoRowsAffected
		return nil, httperror.NewBadRequestErr(ErrUserNotFound, ErrUserNotFound.Error())
	}

	err = s.passwordService.CompareHash(user.Password, data.Password)
	if err != nil {
		return nil, httperror.NewBadRequestErr(ErrUserWrongPassword, ErrUserWrongPassword.Error())
	}

	token, httpError := s.jwtService.GenerateAccessToken(user.ID)
	if httpError != nil {
		return nil, httpError
	}

	return &jwt.JWT{Token: token}, nil
}
