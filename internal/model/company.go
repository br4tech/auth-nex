package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Company struct {
	Id                int      `gorm:"primaryKey"`
	LegalName         string   `gorm:"not null"`
	TradeName         string   `gorm:"not null"`
	Document          string   `gorm:"unique;not null"`
	StateRegistration string   `gorm:"not null"`
	Type              string   // Company type (MEI, ME, LTDA, etc.)
	TenantId          int      `gorm:"column:tenant_id"`
	Schema            string   `gorm:"not null"`
	Address           Address  `gorm:"polymorphic:Addressable;"`
	ParentCompanyId   int      `gorm:"default:null"` // ID da empresa matriz (pode ser nulo)
	ParentCompany     *Company `gorm:"foreignKey:ParentCompanyId"`
	User              []User   `gorm:"many2many:user_companies;"`
	Partners          []Partner
	Activities        []Activity
}

func (model Company) ToDomain() *domain.Company {
	return &domain.Company{
		Id:                model.Id,
		LegalName:         model.LegalName,
		TradeName:         model.TradeName,
		Document:          model.Document,
		StateRegistration: model.StateRegistration,
		Type:              model.Type,
		TenantId:          model.TenantId,
		Schema:            model.Schema,
		ParentCompanyId:   model.ParentCompanyId,
		Address:           *model.Address.ToDomain(),
		Users:             model.convertUsersToDomain(),
		Partners:          model.convertPartnersToDomain(),
		Activities:        model.convertActivitiesToDomain(),
	}
}

func (model *Company) FromDomain(domain *domain.Company) {
	model.Id = domain.Id
	model.LegalName = domain.LegalName
	model.TradeName = domain.TradeName
	model.Document = domain.Document
	model.StateRegistration = domain.StateRegistration
	model.Type = domain.Type
	model.TenantId = int(domain.TenantId)
	model.Schema = domain.Schema
	model.ParentCompanyId = domain.ParentCompanyId
	model.User = convertUsersFromDomain(domain.Users)
	model.Partners = convertPartnersFromDomain(domain.Partners)
	model.Address = Address{
		Street:     domain.Address.Street,
		Number:     domain.Address.Number,
		Complement: domain.Address.Complement,
		District:   domain.Address.District,
		City:       domain.Address.City,
		State:      domain.Address.State,
		ZipCode:    domain.Address.ZipCode,
	}
}

func (model *Company) convertUsersToDomain() []domain.User {
	domainUsers := make([]domain.User, len(model.User))
	for i, user := range model.User {
		domainUsers[i] = *user.ToDomain()
	}
	return domainUsers
}

func (model *Company) convertPartnersToDomain() []domain.Partner {
	domainPartners := make([]domain.Partner, len(model.Partners))
	for i, partner := range model.Partners {
		domainPartners[i] = *partner.ToDomain()
	}
	return domainPartners
}

func (model *Company) convertActivitiesToDomain() []domain.Activity {
	domainActivities := make([]domain.Activity, len(model.Activities))
	for i, activy := range model.Activities {
		domainActivities[i] = *activy.ToDomain()
	}
	return domainActivities
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
