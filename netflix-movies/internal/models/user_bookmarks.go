package models

import "github.com/google/uuid"

type UserMovieBookmark struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"userId"`
	MovieID uuid.UUID `json:"movieId"`
}

// type UserMovieBookmarks []UserMovieBookmark
