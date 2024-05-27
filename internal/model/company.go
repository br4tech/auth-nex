package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	LegalName         string   `gorm:"not null"`
	TradeName         string   `gorm:"not null"`
	Document          string   `gorm:"unique;not null"`
	StateRegistration string   `gorm:"not null"`
	Address           *Address `gorm:"polymorphic:Addressable;"`
	Type              string   // Company type (MEI, ME, LTDA, etc.)
	TenantID          int      `gorm:"column:tenant_id"`
	Schema            string   `gorm:"not null"`
	User              []User   `gorm:"many2many:user_companies;"`
	Partners          []Partner
	Activities        []Activity
}

func (model *Company) ToDomain() *domain.Company {
	companyDomain := &domain.Company{
		LegalName:         model.LegalName,
		TradeName:         model.TradeName,
		Document:          model.Document,
		StateRegistration: model.StateRegistration,
		Type:              model.Type,
		TenantID:          uint(model.TenantID),
		Schema:            model.Schema,
	}

	if model.Address != nil {
		companyDomain.Address = *model.Address.ToDomain()
	}

	if model.Partners != nil {
		companyDomain.Partners = model.convertPartnersToDomain()
	}

	return companyDomain
}

func (model *Company) convertPartnersToDomain() []domain.Partner {
	domainPartners := make([]domain.Partner, len(model.Partners))
	for i, partner := range model.Partners {
		domainPartners[i] = *partner.ToDomain()
	}
	return domainPartners
}

func (model *Company) FromDomain(domain *domain.Company) {
	model.LegalName = domain.LegalName
	model.TradeName = domain.TradeName
	model.Document = domain.Document
	model.StateRegistration = domain.StateRegistration
	model.Address = &Address{
		Street:     domain.Address.Street,
		Number:     domain.Address.Number,
		Complement: domain.Address.Complement,
		District:   domain.Address.District,
		City:       domain.Address.City,
		State:      domain.Address.State,
		ZipCode:    domain.Address.ZipCode,
	}
	model.Type = domain.Type
	model.TenantID = int(domain.TenantID)
	model.Schema = domain.Schema
}
