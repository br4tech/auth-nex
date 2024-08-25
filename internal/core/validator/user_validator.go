package validator

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/go-playground/validator"
)

func ValidateUserSystem(user *domain.User) error {
	validate := validator.New()
	if user.Role != "system" {
		return errors.New("invalid user role")
	}
	if user.Email == "" || user.CPF == "" {
		return errors.New("missing required fields for system user")
	}
	return validate.Struct(user)
}

func ValidateUserClient(user *domain.User) error {
	validate := validator.New()
	if user.Role != "client" {
		return errors.New("invalid user role")
	}
	if user.Email == "" && user.Phone == "" {
		return errors.New("email or phone is required")
	}
	return validate.StructPartial(user, "Name", "Password", "Role", "TenantID", "Email", "Phone")
}
