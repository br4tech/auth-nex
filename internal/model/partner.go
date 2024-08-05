package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type Partner struct {
	Model
	Participation float64 `gorm:"not null"`
	UserId        int
	CompanyId     int
}

func (model *Partner) ToDomain() *domain.Partner {
	return &domain.Partner{
		UserId:        model.UserId,
		Participation: model.Participation,
	}
}

func (model *Partner) FromDomain(domain *domain.Partner) {
	model.Participation = domain.Participation
}
