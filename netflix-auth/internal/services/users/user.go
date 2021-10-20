package users

import (
	"netflix-auth/internal/models"
	"netflix-auth/internal/repository/users"
	httperror "netflix-auth/pkg/http_error"
	"netflix-auth/pkg/passwords"
	"netflix-auth/pkg/postgres"
)

type UserService interface {
	Create(data models.User) (models.User, httperror.HTTPError)
}

type userService struct {
	repo            users.UserRepository
	passwordService passwords.Service
}

func NewUserService(r users.UserRepository, ps passwords.Service) UserService {
	return &userService{
		repo:            r,
		passwordService: ps,
	}
}

func (s userService) Create(input models.User) (models.User, httperror.HTTPError) {
	input.Password = s.passwordService.GeneratePasswordHash(input.Password)
	u, err := s.repo.Create(input)
	if err != nil {
		if postgres.DuplicateError(err) {
			return u, httperror.NewBadRequestErr(ErrUserAlreadyExists, "user with given email already exists")
		}

		return u, httperror.NewInternalServerErr(err)
	}

	return u, nil
}
