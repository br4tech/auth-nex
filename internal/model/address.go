package model

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Street          string `gorm:"not null"`
	Number          string `gorm:"not null"`
	Complement      string
	District        string
	City            string `gorm:"not null"`
	State           string `gorm:"not null"`
	ZipCode         string `gorm:"not null"`
	AddressableID   uint
	AddressableType string
}

func (model *Address) ToDomain() *domain.Address {
	return &domain.Address{
		Street:     model.Street,
		Number:     model.Number,
		Complement: model.Complement,
		District:   model.District,
		City:       model.City,
		State:      model.State,
		ZipCode:    model.ZipCode,
	}
}

func (model *Address) FromDomain(domain *domain.Address) {
	model.Street = domain.Street
	model.Number = domain.Number
	model.Complement = domain.Complement
	model.District = domain.District
	model.City = domain.City
	model.State = domain.State
	model.ZipCode = domain.ZipCode
}
