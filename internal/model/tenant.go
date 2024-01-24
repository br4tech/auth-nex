package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model

	Id        int       `gorm:"primary_key"`
	Name      string    `gorm:unique;not null`
	Companies []Company `gorm:"foreignKey:TenantId"`
	Users     []User    `gorm:"many2many:tenant_admins;"`
}

func (model Tenant) ToDomain() *domain.Tenant {
	domainTenant := &domain.Tenant{
		Name:      model.Name,
		Companies: make([]domain.Company, len(model.Company)),
		Users:     make([]domain.User, len(model.Users)),
	}

	for i, company := range model.Companies {
		domainTenant.Companies[i] = *company.ToDomain()
	}

	for i, user := range model.Users {
		domainTenant.Users[i] = *user.ToDomain()
	}

	return domainTenant
}

func (model Tenant) FromDomain(domain *domain.Tenant) {
	model.Name = domain.Name
	model.Companies = domain.Companies
	model.Users = domain.Users
}
