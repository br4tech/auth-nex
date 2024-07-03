package dto

import "github.com/br4tech/auth-nex/internal/core/domain"

type UserSystemDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	CPF      string `json:"cpf" validate:"required,cpf"`
	Password string `json:"password" validate:"required,password"`
	TenantId int    `json:"tenant_id" validate:"required,tenant_id"`
	Roles    string `json:"roles" validate:"required"`
}

func (dto *UserSystemDTO) ToDomain() *domain.User {
	return &domain.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		TenantId: dto.TenantId,
	}
}
