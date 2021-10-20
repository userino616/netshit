package movies

import (
	"netflix-movies/internal/models"
	"netflix-movies/internal/repository"

	"github.com/google/uuid"
)

type MovieService interface {
	Search(name string) (models.Movies, error)

	GetWatchedList(userID uuid.UUID) (models.Movies, error)
	GetBookmarks(userID uuid.UUID) (models.Movies, error)

	AddToBookmark(data *models.UserMovieBookmark) error
	AddToWatchedList(data *models.UserMovieWatched) error
}

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *movieService {
	return &movieService{repo: repo}
}

func (s *movieService) Search(name string) (models.Movies, error) {
	movies, err := s.repo.Search(name)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *movieService) GetWatchedList(userID uuid.UUID) (models.Movies, error) {
	movies, err := s.repo.GetWatchedList(userID)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *movieService) GetBookmarks(userID uuid.UUID) (models.Movies, error) {
	movies, err := s.repo.GetBookmarks(userID)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *movieService) AddToBookmark(data *models.UserMovieBookmark) error {
	err := s.repo.AddToBookmark(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *movieService) AddToWatchedList(data *models.UserMovieWatched) error {
	err := s.repo.AddToWatchedList(data)
	if err != nil {
		return err
	}

	return nil
}
