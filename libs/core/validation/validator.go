package validation

import (
	"libs/core/validation/http"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func RegisterCustomHttpValidators(validator *validator.Validate) (err error) {
	if err = validator.RegisterValidation("http_method", http.ValidateHttpMethod); err != nil {
		return
	}
	return
}
