package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"unique;not null"`
	Users []User `gorm:"many2many:user_roles"`
}
