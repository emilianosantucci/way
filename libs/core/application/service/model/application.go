package model

import (
	"github.com/google/uuid"
)

type Application struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid4_rfc4122"`
	Name    string    `json:"name" validate:"required,min=3,max=50"`
	Version string    `json:"version" validate:"required,min=1,max=50"`
}
