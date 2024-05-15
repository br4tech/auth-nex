package dto

type TenantDTO struct {
	Name string `json:"name" validate:"required"`
}
