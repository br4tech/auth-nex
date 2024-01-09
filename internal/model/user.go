package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm:"unique;not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Role     []Role `gorm:"many2many:user_roles;"`
	TenantID int    `gorm:"column:tenant_id"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func (model User) ToDomain() *domain.User {
	return &domain.User{
		Name: model.Name,
	}
}

func (model User) FromDomain(entity *domain.User) {
	model.Name = entity.Name
}
