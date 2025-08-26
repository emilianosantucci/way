package dto

import (
	"github.com/google/uuid"
)

type Route struct {
	ID          uuid.UUID `json:"id"`
	Path        string    `json:"path"`
	Name        string    `json:"name"`
	Component   string    `json:"component"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
}
