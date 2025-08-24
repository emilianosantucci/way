package model

import (
	"libs/core/common/http"
)

type NewRestApiResource struct {
	Path   string          `json:"path" validate:"required,min=1"`
	Method http.HttpMethod `json:"method" validate:"http_method"`
}
