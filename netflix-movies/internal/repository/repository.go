package repository

import (
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	MovieRepository
}

func New(db *pg.DB) *Repository {
	return &Repository{
		MovieRepository: NewMovieRepository(db),
	}
}
