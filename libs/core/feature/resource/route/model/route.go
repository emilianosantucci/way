package model

import (
	"github.com/google/uuid"
)

type Route struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid4_rfc4122"`
	Path        string    `json:"path" validate:"required,min=1"`
	Name        string    `json:"name" validate:"required,min=1"`
	Description string    `json:"description" validate:"omitempty"`
}
