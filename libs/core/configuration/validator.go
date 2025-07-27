package configuration

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

type ValidatorResult struct {
	fx.Out
	Validator *validator.Validate
}

func NewValidator() ValidatorResult {
	return ValidatorResult{Validator: validator.New(validator.WithRequiredStructEnabled())}
}
