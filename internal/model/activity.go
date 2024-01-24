package model

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	CNAE        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CompanyID   uint   `gorm:"column:comapany_id"`
}
