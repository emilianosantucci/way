package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestApiResource struct {
	gorm.Model
	ID     uuid.UUID `json:"id"`
	Path   string    `json:"path"`
	Method string    `json:"method"`
}
