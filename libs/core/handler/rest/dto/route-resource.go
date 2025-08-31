package dto

import (
	"github.com/google/uuid"
)

type RouteResource struct {
	ID          uuid.UUID `json:"id"`
	Path        string    `json:"path"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
