package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Model
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CPF       string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Phone     string
	TenantId  int `gorm:"unique;not null"`
	Tenant    Tenant
	Role      string `gorm:"not null"`
	ProfileId int    `gorm:"unique;not null"`
	Profile   Profile
	Companies []Company `gorm:"many2many:user_companies;"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (model User) ToDomain() *domain.User {
	return &domain.User{
		Name:      model.Name,
		Email:     model.Email,
		CPF:       &model.CPF,
		Password:  model.Password,
		Phone:     model.Phone,
		TenantId:  model.TenantId,
		Role:      model.Role,
		ProfileId: &model.ProfileId,
	}
}

func (model *User) FromDomain(domain *domain.User) {
	model.Name = domain.Name
	model.Email = domain.Email
	model.CPF = *domain.CPF
	model.Password = domain.Password
	model.Phone = domain.Phone
	model.TenantId = domain.TenantId
	model.Role = domain.Role
	model.ProfileId = *domain.ProfileId
}
