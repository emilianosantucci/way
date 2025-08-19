package dto

type NewApplication struct {
	Name    string `json:"name" validate:"required,min=3,max=50"`
	Version string `json:"version" validate:"required,min=1,max=50"`
}
