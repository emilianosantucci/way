package model

import "github.com/google/uuid"

type UpdateApplication struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid4_rfc4122"`
	Version string    `json:"version" validate:"omitempty,min=1,max=50"`
}
