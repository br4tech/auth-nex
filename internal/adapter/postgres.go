package adapter

import (
	"errors"
	"fmt"

	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/core/port"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAdapter[T port.IModel] struct {
	Db *gorm.DB
}

func NewPostgresAdapter[T port.IModel](cfg *config.Config) port.IDatabase[T] {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	return &PostgresAdapter[T]{Db: db}
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

func (adapter *PostgresAdapter[T]) FindById(id int) (*T, error) {
	if adapter.Db == nil {
		return nil, errors.New("database connection not established")
	}

	var result T
	if err := adapter.Db.First(&result, id).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (adapter *PostgresAdapter[T]) FindBy(field string, value string) ([]*T, error) {
	if adapter.Db == nil {
		return nil, errors.New("database connection not established")
	}

	var results []*T
	if err := adapter.Db.Where(field+"= ?", value).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (adapter *PostgresAdapter[T]) Create(entity T) (int, error) {
	if adapter.Db == nil {
		return 0, errors.New("database connection not established")
	}

	create := adapter.Db.Create(entity)
	if create.Error != nil {
		return 0, create.Error
	}

	id := entity.GetId()

	return id, nil
}
