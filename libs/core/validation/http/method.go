package http

import (
	"libs/core/common/http"

	"github.com/go-playground/validator/v10"
)

func ValidateHttpMethod(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(http.HttpMethod)
	return value != http.Unknown
}
