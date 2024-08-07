package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Profile struct {
	Model

	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_profiles"`
}

func (model Profile) GetId() int {
	return model.Id
}

func (model *Profile) ToDomain() *domain.Profile {
	return &domain.Profile{
		Name: model.Name,
	}
}

func (model *Profile) FromDomain(domain *domain.Profile) {
	model.Name = domain.Name
}
