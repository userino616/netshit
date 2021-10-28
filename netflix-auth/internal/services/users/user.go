package users

import (
	"errors"
	"netflix-auth/internal/models"
	"netflix-auth/internal/repository/users"
	"netflix-auth/pkg/hash"
	httperror "netflix-auth/pkg/http_error"
	"netflix-auth/pkg/postgres"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type Service interface {
	Create(data models.User) (models.User, httperror.HTTPError)
	GetByEmail(email string) (*models.User, httperror.HTTPError)
	GetByID(id uuid.UUID) (*models.User, httperror.HTTPError)
}

type userService struct {
	repo users.UserRepository
	hs   hash.Service
}

func NewUserService(r users.UserRepository, hs hash.Service) Service {
	return &userService{
		repo: r,
		hs:   hs,
	}
}

func (s userService) Create(input models.User) (models.User, httperror.HTTPError) {
	input.Password = s.hs.Generate(input.Password)
	u, err := s.repo.Create(input)
	if err != nil {
		if postgres.DuplicateError(err) {
			return u, httperror.NewBadRequestErr(ErrUserAlreadyExists, "user with given email already exists")
		}

		return u, httperror.NewInternalServerErr(err)
	}

	return u, nil
}

func (s userService) GetByID(id uuid.UUID) (*models.User, httperror.HTTPError) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, httperror.NewNotFoundErr(err, "user not found")
		}

		return nil, httperror.NewInternalServerErr(err)
	}

	return &user, nil
}

func (s userService) GetByEmail(email string) (*models.User, httperror.HTTPError) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, httperror.NewNotFoundErr(err, "user not found")
		}

		return nil, httperror.NewInternalServerErr(err)
	}

	return &user, nil
}
