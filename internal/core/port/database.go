package port

import "gorm.io/gorm"

type IDatabase[T IModel] interface {
	GetDb() *gorm.DB
	FindAll() ([]T, error)
	Create(entity T) (int, error)
}

type IModel interface {
	GetId() int
}
