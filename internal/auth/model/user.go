package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       int    `gorm:"primary:key"`
	UserName string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     []Role `gorm:"many2many:user_permissions;"`
	TenantID int    `gorm:"column:tenant_id"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
