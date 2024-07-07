package port

import "gorm.io/gorm"

type IDatabase interface {
	GetDb() *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) (*gorm.DB, error)
	Error() error
}
