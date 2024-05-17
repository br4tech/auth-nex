package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model

	Id        int       `gorm:"primary_key;"`
	Name      string    `gorm:"unique;not null;"`
	Companies []Company `gorm:"many2many:tenant_companies;"`
	Users     []User    `gorm:"many2many:tenant_users;"`
}

func (model Tenant) ToDomain() *domain.Tenant {
	return &domain.Tenant{
		Name:      model.Name,
		Companies: model.convertCompaniesToDomain(),
		Users:     model.convertUsersToDomain(),
	}
}

func (model *Tenant) convertCompaniesToDomain() []domain.Company {
	domainCompanies := make([]domain.Company, len(model.Companies))
	for i, company := range model.Companies {
		domainCompanies[i] = *company.ToDomain()
	}
	return domainCompanies
}

func (model *Tenant) convertUsersToDomain() []domain.User {
	domainUsers := make([]domain.User, len(model.Users))
	for i, user := range model.Users {
		domainUsers[i] = *user.ToDomain()
	}
	return domainUsers
}

func (model *Tenant) FromDomain(domain *domain.Tenant) {
	model.Name = domain.Name
	model.Companies = convertCompaniesFromDomain(domain.Companies, model.Id)
	model.Users = convertUsersFromDomain(domain.Users)
}

func convertCompaniesFromDomain(domainCompanies []domain.Company, tenantId int) []Company {
	modelCompanies := make([]Company, len(domainCompanies))
	for i, domainCompany := range domainCompanies {
		modelCompanies[i] = Company{
			LegalName: domainCompany.LegalName,
			TradeName: domainCompany.TradeName,
			TenantID:  tenantId,
		}
	}
	return modelCompanies
}

func convertUsersFromDomain(domainUsers []domain.User) []User {
	modelUsers := make([]User, len(domainUsers))
	for i, domainUser := range domainUsers {
		modelUsers[i] = User{
			Name:  domainUser.Name,
			Email: domainUser.Email,
		}
	}
	return modelUsers
}
