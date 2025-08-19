package dto

type UpdateApplication struct {
	ID      string `json:"id" validate:"required,uuid4_rfc4122"`
	Version string `json:"version" validate:"required,min=1,max=50"`
}
