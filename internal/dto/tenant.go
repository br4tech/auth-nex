package dto

import "github.com/br4tech/auth-nex/internal/core/domain"

type TenantDTO struct {
	Name    string        `json:"name" validate:"required"`
	Company CompanyDTO    `json:"company"`
	User    UserSystemDTO `json:"user"`
}

func (dto *TenantDTO) ToDomain() *domain.Tenant {
	return &domain.Tenant{
		Name: dto.Name,
		Companies: []domain.Company{
			{
				LegalName:         dto.Company.LegalName,
				TradeName:         dto.Company.TradeName,
				Document:          dto.Company.Document,
				StateRegistration: dto.Company.StateRegistration,
				Type:              dto.Company.Type,
			},
		},
		Users: []domain.User{
			{
				Name:  dto.User.Name,
				Email: dto.User.Email,
				CPF:   dto.User.CPF,
				Role:  dto.User.Role,
			},
		},
	}
}

func (dto *TenantDTO) FromDomain(domain *domain.Tenant) {
	dto.Name = domain.Name

	dto.Company.LegalName = domain.Companies[0].LegalName
	dto.Company.TradeName = domain.Companies[0].TradeName
	dto.Company.Document = domain.Companies[0].Document
	dto.Company.StateRegistration = domain.Companies[0].StateRegistration
	dto.Company.Type = domain.Companies[0].Type

	dto.User.Name = domain.Users[0].Name
	dto.User.Email = domain.Users[0].Email
	dto.User.CPF = domain.Users[0].CPF
	dto.User.Role = domain.Users[0].Role
}
