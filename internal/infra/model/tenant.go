package model

import "gorm.io/gorm"

type Tenant struct {
	gorm.Model

	Id   int    `gorm:"primary_key"`
	Name string `gorm:unique;not null`
}
