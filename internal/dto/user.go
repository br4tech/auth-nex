package dto

type UserDTO struct {
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,email"`
	TenantID int      `json:"tenant_id" validate:"required,email"`
	Roles    []string `json:"roles" validate:"required"`
}
