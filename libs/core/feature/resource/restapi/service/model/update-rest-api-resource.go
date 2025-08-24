package model

import (
	"libs/core/common/http"

	"github.com/google/uuid"
)

type UpdateRestApiResource struct {
	ID     uuid.UUID       `json:"id" validate:"required,uuid4_rfc4122"`
	Path   string          `json:"path" validate:"omitempty,min=1"`
	Method http.HttpMethod `json:"method" validate:"omitempty,http_method"`
}
