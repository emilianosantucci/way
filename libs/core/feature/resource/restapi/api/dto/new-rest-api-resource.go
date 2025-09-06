package dto

type NewRestApiResource struct {
	Path   string `json:"path" validate:"required,min=1"`
	Method string `json:"method"`
}
