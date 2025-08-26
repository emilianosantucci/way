package model

type NewRoute struct {
	Path        string `json:"path" validate:"required,min=1"`
	Name        string `json:"name" validate:"required,min=1"`
	Component   string `json:"component" validate:"required,min=1"`
	Description string `json:"description" validate:"omitempty"`
	IsActive    bool   `json:"is_active"`
}
