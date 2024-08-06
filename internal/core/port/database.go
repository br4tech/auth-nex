package port

import "gorm.io/gorm"

type IDatabase[T IModel] interface {
	GetDb() *gorm.DB
	FindAll() ([]T, error)
	FindById(id int) (*T, error)
	FindBy(field string, value string) ([]*T, error)
	Create(entity T) (int, error)
}

type IModel interface {
	GetId() int
}
