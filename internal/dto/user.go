package dto

type UserDTO struct {
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,password"`
	TenantID int      `json:"tenant_id" validate:"required,tenant_id"`
	Roles    []string `json:"roles" validate:"required"`
}
