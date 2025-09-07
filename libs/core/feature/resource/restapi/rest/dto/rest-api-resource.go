package dto

import (
	"github.com/google/uuid"
)

type RestApiResource struct {
	ID     uuid.UUID `json:"id"`
	Path   string    `json:"path"`
	Method string    `json:"method"`
}
