package model

import (
	"github.com/google/uuid"
)

type Application struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Version string    `json:"version"`
}
