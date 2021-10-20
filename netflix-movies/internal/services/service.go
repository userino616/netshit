package services

import (
	"netflix-movies/internal/repository"
	"netflix-movies/internal/services/movies"
)

type Service struct {
	movies.MovieService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		MovieService: movies.NewMovieService(repos.MovieRepository),
	}
}
