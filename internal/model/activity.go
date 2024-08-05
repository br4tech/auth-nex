package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Activity struct {
	Model
	CNAE        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CompanyID   uint   `gorm:"column:comapany_id"`
}

func (model *Activity) ToDomain() *domain.Activity {
	return &domain.Activity{
		CNAE:        model.CNAE,
		Description: model.Description,
		CompanyID:   model.CompanyID,
	}
}

func (model *Activity) FromDomain(domain *domain.Activity) {
	model.CNAE = domain.CNAE
	model.Description = domain.Description
	model.CompanyID = domain.CompanyID
}
