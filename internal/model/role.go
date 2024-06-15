package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_roles"`
}

func (model *Role) ToDomain() *domain.Role {
	return &domain.Role{
		Name: model.Name,
	}
}

func (model *Role) FromDomain(domain *domain.Role) {
	model.Name = domain.Name
}
