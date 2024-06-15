package domain

import (
	"github.com/go-playground/validator"
)

func IsRequiredFieldMissing(err error) bool {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return false
	}

	for _, e := range err.(validator.ValidationErrors) {
		if e.Tag() == "required" {
			return true
		}
	}
	return false
}
