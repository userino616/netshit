package repository

import (
	"netflix-auth/internal/repository/users"

	"github.com/go-pg/pg/v10"
)

type Repository struct {
	User users.UserRepository
}

func New(db *pg.DB) *Repository {
	return &Repository{
		User: users.NewUserRepository(db),
	}
}
