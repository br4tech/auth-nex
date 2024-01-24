package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Participation float64 `gorm:"not null"`
	User          User
}

func (model Partner) ToDomain() *domain.Partner {
	return &domain.Partner{
		Participation: model.Participation,
		User:          model.ToDomain().User,
	}
}
