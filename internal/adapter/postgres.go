package adapter

import (
	"database/sql"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAdapter[T any] struct {
	Db *gorm.DB
}

func NewPostgresAdapter[T any](db *sql.DB) *PostgresAdapter[T] {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &PostgresAdapter[T]{
		Db: gormDB,
	}
}

func (adapter *PostgresAdapter[T]) GetDb() *gorm.DB {
	return adapter.Db
}

func (adapter *PostgresAdapter[T]) FindAll() ([]T, error) {
	if adapter.Db == nil {
		return nil, errors.New("database connection not established")
	}

	var results []T

	result := adapter.Db.Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	return results, nil
}

func (adapter *PostgresAdapter[T]) Find(query interface{}, args ...interface{}) ([]*T, error) {
	if adapter.Db == nil {
		return nil, errors.New("database connection not established")
	}

	var results []*T
	if err := adapter.Db.Where(query, args...).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (adapter *PostgresAdapter[T]) FindBy(query interface{}, args ...interface{}) (*T, error) {
	if adapter.Db == nil {
		return nil, errors.New("database connection not established")
	}

	var result T
	if err := adapter.Db.Where(query, args...).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (adapter *PostgresAdapter[T]) Create(entity T) (int, error) {
	if adapter.Db == nil {
		return 0, errors.New("database connection not established")
	}

	create := adapter.Db.Create(entity)
	if create.Error != nil {
		return 0, create.Error
	}

	return 0, nil
}
