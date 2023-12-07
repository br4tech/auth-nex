package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       int    `gorm:"primary:key"`
	UserName string `gorm:"unique;not null`
	Password string `gorm:"not null"`
	Roles    []string
	TenantID int
}
