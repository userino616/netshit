package repository

import (
	"netflix-auth/internal/repository/users"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
)

type Repository struct {
	User     users.UserRepository
	MemStore MemStore
}

func New(db *pg.DB, rc *redis.Client) *Repository {
	return &Repository{
		User:     users.NewUserRepository(db),
		MemStore: NewMemoryStorage(rc),
	}
}
