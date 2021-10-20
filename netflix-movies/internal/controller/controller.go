package controller

import (
	"netflix-movies/internal/services"

	"github.com/userino616/netflix-grpc/movieservice"
)

type Movie interface {
	movieservice.MovieServiceServer
}

type Controller struct {
	Movie
}

func New(services *services.Service) *Controller {
	return &Controller{
		Movie: NewMovieController(services.MovieService),
	}
}
