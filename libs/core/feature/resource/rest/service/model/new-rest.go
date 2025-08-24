package model

import (
	"libs/core/common/http"
)

type NewRestResource struct {
	Path   string          `json:"path" validate:"required,min=1"`
	Method http.HttpMethod `json:"method" validate:"omitnil,http_method"`
}
