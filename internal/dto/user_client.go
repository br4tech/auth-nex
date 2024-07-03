package dto

import "github.com/br4tech/auth-nex/internal/core/domain"

type UserClientDTO struct {
	Phone    string `json:"phone"`
	TenantId int    `json:"tenant_id" validate:"required,tenant_id"`
	Roles    string `json:"roles" validate:"required"`
}

func (dto *UserClientDTO) ToDomain() *domain.User {
	return &domain.User{
		Phone:    dto.Phone,
		TenantId: dto.TenantId,
	}
}
