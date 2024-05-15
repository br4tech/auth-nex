package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name          string  `gorm:"not null"`
	Email         string  `gorm:"unique;not null"`
	Password      string  `gorm:"not null"`
	CPF           string  `gorm:unique;not null`
	Nationality   string  `gorm:"not null"`
	MaritalStatus string  `gorm:"not null"`
	Address       Address `gorm:"polymorphic:Addressable;"`
	TenantId      int
	Tenant        Tenant
	Companies     []Company `gorm:"many2many:user_companies;"`
	Role          []Role    `gorm:"many2many:user_roles;"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (model *User) ToDomain() *domain.User {
	return &domain.User{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		TenantId: model.TenantId,
		Roles:    []domain.Role{},
	}
}

func (model *User) FromDomain(domain *domain.User) {
	model.Name = domain.Name
	model.Email = domain.Email
	model.Password = domain.Password
	model.TenantId = domain.TenantId
}
