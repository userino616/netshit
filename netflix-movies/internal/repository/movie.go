package repository

import (
	"fmt"
	"netflix-movies/internal/models"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type MovieRepository interface {
	Search(name string) (models.Movies, error)

	GetWatchedList(userID uuid.UUID) (models.Movies, error)
	GetBookmarks(userID uuid.UUID) (models.Movies, error)

	AddToBookmark(data *models.UserMovieBookmark) error
	AddToWatchedList(data *models.UserMovieWatched) error
}

type movieRepository struct {
	db *pg.DB
}

func NewMovieRepository(db *pg.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r movieRepository) Search(name string) (movies models.Movies, err error) {
	err = r.db.
		Model(&movies).
		Where("lower(name) LIKE ?", fmt.Sprintf("%%%s%%", strings.ToLower(name))).
		Select()

	return
}

func (r movieRepository) AddToBookmark(data *models.UserMovieBookmark) (err error) {
	_, err = r.db.
		Model(data).
		Insert()

	return
}

func (r movieRepository) AddToWatchedList(data *models.UserMovieWatched) (err error) {
	// _, err = r.db.
	//	Model(data).
	//	Table("user_movie_watched").
	//	Insert()

	// #42P01 relation "user_movie_watcheds" does not exist ??
	query := "INSERT INTO user_movie_watched (user_id, movie_id)\nVALUES (?, ?)"
	_, err = r.db.ExecOne(query, data.UserID, data.MovieID)

	return
}

func (r movieRepository) GetWatchedList(userID uuid.UUID) (movies models.Movies, err error) {
	query := "SELECT * FROM movies WHERE id IN (SELECT movie_id FROM user_movie_watched WHERE user_id = ?)"
	_, err = r.db.Query(&movies, query, userID)

	return
}

func (r movieRepository) GetBookmarks(userID uuid.UUID) (movies models.Movies, err error) {
	query := "SELECT * FROM movies WHERE id IN (SELECT movie_id FROM user_movie_bookmarks WHERE user_id = ?)"
	_, err = r.db.Query(&movies, query, userID)

	return
}
