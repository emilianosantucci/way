package dto

type UpdateRoute struct {
	ID          string `json:"id" validate:"required,uuid4_rfc4122"`
	Path        string `json:"path" validate:"omitempty,min=1"`
	Name        string `json:"name" validate:"omitempty,min=1"`
	Component   string `json:"component" validate:"omitempty,min=1"`
	Description string `json:"description" validate:"omitempty"`
	IsActive    *bool  `json:"is_active" validate:"omitempty"`
}
