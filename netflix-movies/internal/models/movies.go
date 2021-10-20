package models

import "github.com/google/uuid"

type Movie struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Movies []Movie
