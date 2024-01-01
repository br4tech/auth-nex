package port

import "gorm.io/gorm"

type IDatabase interface {
	GetDb() *gorm.DB
}
