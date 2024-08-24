package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Activity struct {
	Id          int    `gorm:"primaryKey"`
	CNAE        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (model *Activity) ToDomain() *domain.Activity {
	return &domain.Activity{
		CNAE:        model.CNAE,
		Description: model.Description,
	}
}

func (model *Activity) FromDomain(domain *domain.Activity) {
	model.CNAE = domain.CNAE
	model.Description = domain.Description
}
