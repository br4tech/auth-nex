package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Profile struct {
	Id          int          `gorm:"primaryKey"`
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:profile_permissions;"`
}

func (model *Profile) ToDomain() *domain.Profile {
	return &domain.Profile{
		Name: model.Name,
	}
}

func (model *Profile) FromDomain(domain *domain.Profile) {
	model.Name = domain.Name
}
