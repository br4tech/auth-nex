package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Company struct {
	Model
	LegalName         string  `gorm:"not null"`
	TradeName         string  `gorm:"not null"`
	Document          string  `gorm:"unique;not null"`
	StateRegistration string  `gorm:"not null"`
	Address           Address `gorm:"polymorphic:Addressable;"`
	Type              string  // Company type (MEI, ME, LTDA, etc.)
	TenantId          int     `gorm:"column:tenant_id"`
	Schema            string  `gorm:"not null"`
	User              []User  `gorm:"many2many:user_companies;"`
	Partners          []Partner
	Activities        []Activity
}

func (model Company) GetId() int {
	return model.Id
}

func (model Company) ToDomain() *domain.Company {
	companyDomain := &domain.Company{
		Id:                model.Id,
		LegalName:         model.LegalName,
		TradeName:         model.TradeName,
		Document:          model.Document,
		StateRegistration: model.StateRegistration,
		Type:              model.Type,
		TenantId:          uint(model.TenantId),
		Schema:            model.Schema,
	}

	companyDomain.Address = *model.Address.ToDomain()

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
	model.Id = domain.Id
	model.LegalName = domain.LegalName
	model.TradeName = domain.TradeName
	model.Document = domain.Document
	model.StateRegistration = domain.StateRegistration
	model.Address = Address{
		Street:     domain.Address.Street,
		Number:     domain.Address.Number,
		Complement: domain.Address.Complement,
		District:   domain.Address.District,
		City:       domain.Address.City,
		State:      domain.Address.State,
		ZipCode:    domain.Address.ZipCode,
	}
	model.Type = domain.Type
	model.TenantId = int(domain.TenantId)
	model.Schema = domain.Schema
	model.Partners = convertPartnersFromDomain(domain.Partners)
}

func convertPartnersFromDomain(domainPartners []domain.Partner) []Partner {
	modelPartnes := make([]Partner, len(domainPartners))
	for i, domainPartner := range domainPartners {
		modelPartnes[i] = Partner{
			Participation: domainPartner.Participation,
			UserId:        domainPartner.UserId,
			CompanyId:     domainPartner.CompanyId,
		}
	}
	return modelPartnes
}
