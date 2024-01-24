package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	LegalName         string `gorm:"not null"`
	TradeName         string `gorm:"not null"`
	CNPJ              string `gorm:unique;not null`
	StateRegistration string `gorm:"not null"`
	Address           Address
	Partners          []Partner
	Activities        []Activity
	Type              string // Company type (MEI, ME, LTDA, etc.)
	TenantID          uint   `gorm:"column:tenant_id"`
	Schema            string `gorm:"not null"`
	User              []User
}

func (model Company) ToDomain() *domain.Company {
	return &domain.Company{
		LegalName:         model.LegalName,
		TradeName:         model.TradeName,
		CNPJ:              model.CNPJ,
		StateRegistration: model.StateRegistration,
		Address:           *model.Address.ToDomain(),
		Partners:          model.Partners.ToDomain(),
	}
}

func (model *Company) convertPartners(domainPartners []domain.Partner) []Partner {
	partners := make([]Partner, len(domainPartners))
	for i, domainPartner := range domainPartners {
		partners[i] = *domainPartner.FromDomain()
	}
	return partners
}
