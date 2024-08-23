package validator

import (
	"fmt"

	"github.com/go-playground/validator"
)

type DomainErrorType byte

const (
	INVALID_STRUCT_FIELDS DomainErrorType = 0
	ENTITY_NOT_FOUND      DomainErrorType = 1
)

type DomainError struct {
	err       string
	errorType DomainErrorType
}

func (e DomainError) Error() string {
	return e.err
}

func (e DomainError) Type() DomainErrorType {
	return e.errorType
}

func NewDomainError(errorType DomainErrorType, text string) DomainError {
	return DomainError{
		err:       text,
		errorType: errorType,
	}
}

func ValidateStruct(d any) error {
	validate := validator.New()
	err := validate.Struct(d)

	if err != nil {
		message := ""
		for _, e := range err.(validator.ValidationErrors) {
			message += fmt.Sprintf("%s: %s\n", e.Field(), e.Tag())
		}

		return NewDomainError(INVALID_STRUCT_FIELDS, message)
	}

	return nil
}
