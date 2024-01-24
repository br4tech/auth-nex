package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name          string `gorm:"not null"`
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	CPF           string `gorm:unique;not null`
	Nationality   string `gorm:"not null"`
	MaritalStatus string `gorm:"not null"`
	Address       Address
	Role          []Role    `gorm:"many2many:user_roles;"`
	TenantID      int       `gorm:"column:tenant_id"`
	Companies     []Company `gorm:"many2many:user_companies;"`
}

type Doctor struct {
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (model User) ToDomain() *domain.User {
	return &domain.User{
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		TenantID: model.TenantID,
	}
}

func (model *User) FromDomain(domain *domain.User) {
	model.Name = domain.Name
	model.Email = domain.Email
	model.Password = domain.Password
	model.TenantID = domain.TenantID
}
