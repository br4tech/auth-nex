package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Tenant struct {
	gorm.Model

	Id   int    `gorm:"primary_key"`
	Name string `gorm:unique;not null`
}

func (model Tenant) ToDomain() *domain.Tenant {
	return &domain.Tenant{
		Name: model.Name,
	}
}

func (model Tenant) FromDomain(entity *domain.Tenant) {
	model.Name = entity.Name
}
