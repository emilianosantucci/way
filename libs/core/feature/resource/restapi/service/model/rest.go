package model

import (
	"libs/core/common/http"

	"github.com/google/uuid"
)

type RestApiResource struct {
	ID     uuid.UUID       `json:"id" validate:"required,uuid4_rfc4122"`
	Path   string          `json:"path" validate:"required,min=1"`
	Method http.HttpMethod `json:"method" validate:"http_method"`
}
