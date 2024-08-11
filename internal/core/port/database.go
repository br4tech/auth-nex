package port

import "gorm.io/gorm"

type IDatabase[T IModel] interface {
	GetDb() *gorm.DB
	FindAll() ([]T, error)
	Find(query interface{}, args ...interface{}) ([]*T, error)
	FindBy(query interface{}, args ...interface{}) (*T, error)
	Create(entity T) (int, error)
}

type IModel interface {
	GetId() int
}
