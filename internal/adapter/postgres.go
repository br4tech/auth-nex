package adapter

import (
	"context"
	"fmt"

	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/core/port"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) port.IDatabase[any] {
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

	return &PostgresDatabase{Db: db}
}

func (p *PostgresDatabase) FindOne(ctx context.Context, conds ...interface{}) (any, error) {
	var result any

	query := p.Db.WithContext(ctx).Model(&result)
	if err := query.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (p *PostgresDatabase) Create(ctx context.Context, value ...interface{}) (*any, error) { // Alterado para retornar *any
	db := p.Db

	if err := db.WithContext(ctx).Create(value).Error; err != nil {
		return nil, err
	}

	return value, nil
}
