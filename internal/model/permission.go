package model

import "github.com/br4tech/auth-nex/internal/core/domain"

type Permission struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
}

func (model *Permission) ToDomain() *domain.Permission {
	return &domain.Permission{
		Id:   model.Id,
		Name: model.Name,
	}
}

func (model *Permission) FromDomain(domain *domain.Permission) {
	model.Id = domain.Id
	model.Name = domain.Name
}
