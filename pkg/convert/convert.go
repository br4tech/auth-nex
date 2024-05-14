package convert

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/model"
)

func (model *model.Tenant) ConvertCompaniesToDomain() []domain.Company {
	domainCompanies := make([]domain.Company, len(model.Companies))
	for i, company := range model.Companies {
		domainCompanies[i] = *company.ToDomain()
	}
	return domainCompanies
}

func (model *model.Tenant) ConvertUsersToDomain() []domain.User {
	domainUsers := make([]domain.User, len(model.Users))
	for i, user := range model.Users {
		domainUsers[i] = *user.ToDomain()
	}
	return domainUsers
}

func ConvertCompaniesFromDomain(domainCompanies []domain.Company, tenantId int) []model.Company {
	modelCompanies := make([]model.Company, len(domainCompanies))
	for i, domainCompany := range domainCompanies {
		modelCompanies[i] = model.Company{
			LegalName: domainCompany.LegalName,
			TradeName: domainCompany.TradeName,
			TenantID:  tenantId,
		}
	}
	return modelCompanies
}

func ConvertUsersFromDomain(domainUsers []domain.User) []model.User {
	modelUsers := make([]model.User, len(domainUsers))
	for i, domainUser := range domainUsers {
		modelUsers[i] = model.User{
			Name:  domainUser.Name,
			Email: domainUser.Email,
		}
	}
	return modelUsers
}
