package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	LegalName         string  `gorm:"not null"`
	TradeName         string  `gorm:"not null"`
	CNPJ              string  `gorm:unique;not null`
	StateRegistration string  `gorm:"not null"`
	Address           Address `gorm:"polymorphic:Addressable;"`
	Type              string  // Company type (MEI, ME, LTDA, etc.)
	TenantID          int     `gorm:"column:tenant_id"`
	Schema            string  `gorm:"not null"`
	User              []User  `gorm:"many2many:user_companies;"`
	Partners          []Partner
	Activities        []Activity
	Tenants           []Tenant `gorm:"many2many:tenant_companies;"`
}

func (model Company) ToDomain() *domain.Company {
	return &domain.Company{
		LegalName:         model.LegalName,
		TradeName:         model.TradeName,
		CNPJ:              model.CNPJ,
		StateRegistration: model.StateRegistration,
		Address:           *model.Address.ToDomain(),
		Partners:          model.convertPartnersToDomain(),
	}
}

func (model *Company) convertPartnersToDomain() []domain.Partner {
	domainPartners := make([]domain.Partner, len(model.Partners))
	for i, partner := range model.Partners {
		domainPartners[i] = *partner.ToDomain()
	}
	return domainPartners
}
