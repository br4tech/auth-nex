package model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserName string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     []Role `gorm:"many2many:user_roles;"`
	TenantID int    `gorm:"column:tenant_id"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
