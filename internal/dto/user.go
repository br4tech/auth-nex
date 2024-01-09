package dto

type UserDTO struct {
	Name  string   `json:"name" validate:"required"`
	Email string   `json:"email" validate:"required,email"`
	Roles []string `json:"roles" validate:"required"`
}
