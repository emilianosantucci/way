package dto

type UpdateRestApiResource struct {
	ID     string `json:"id" validate:"required,uuid4_rfc4122"`
	Path   string `json:"path" validate:"omitempty,min=1"`
	Method string `json:"method" validate:"omitempty"`
}
