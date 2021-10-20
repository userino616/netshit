package models

import "github.com/google/uuid"

type UserMovieWatched struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	MovieID uuid.UUID `json:"movieId"`

	tableName struct{} `pg:"user_movie_watched"`
}

type UserMovieWatchedList []UserMovieWatched
