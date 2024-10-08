package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string
	CPF      string
	Password string `gorm:"not null"`
	Phone    string
	Role     string `gorm:"not null"`

	TenantId int
	Tenant   Tenant `gorm:"foreignKey:TenantId"`

	Profiles []Profile `gorm:"many2many:user_profiles;"`
}

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func (model *User) ToDomain() *domain.User {
	var domainProfiles []domain.Profile
	for _, profileModel := range model.Profiles {
		domainProfiles = append(domainProfiles, *profileModel.ToDomain())
	}

	return &domain.User{
		Name:     model.Name,
		Email:    model.Email,
		CPF:      model.CPF,
		Password: model.Password,
		Phone:    model.Phone,
		TenantId: model.TenantId,
		Role:     model.Role,
		Profiles: domainProfiles,
	}
}

func (model *User) FromDomain(domain *domain.User) {
	model.Name = domain.Name
	model.Email = domain.Email
	model.CPF = domain.CPF
	model.Password = domain.Password
	model.Phone = domain.Phone
	model.TenantId = domain.TenantId
	model.Role = domain.Role

	var profileModels []Profile
	for _, domainProfile := range domain.Profiles {
		profileModel := Profile{}
		profileModel.FromDomain(&domainProfile)
		profileModels = append(profileModels, profileModel)
	}

	model.Profiles = profileModels
}
