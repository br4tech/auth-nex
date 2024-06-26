package dto

type TenantDTO struct {
	Name    string     `json:"name" validate:"required"`
	Company CompanyDTO `json:"company"`
	User    UserDTO    `json:"user"`
}
