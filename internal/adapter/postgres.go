package adapter

import (
	"fmt"

	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/core/port"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) port.IDatabase {
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

	return &postgresDatabase{Db: db}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}

func (p *postgresDatabase) Where(query interface{}, args ...interface{}) *gorm.DB {
	return p.Db.Where(query, args...)
}

func (p *postgresDatabase) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return p.Db.First(dest, conds...)
}

func (p *postgresDatabase) Create(value interface{}) (*gorm.DB, error) {
	create := p.Db.Create(value)
	if create.Error != nil {
		return nil, create.Error
	}
	return create, nil
}

func (p *postgresDatabase) Error() error {
	return p.Db.Error
}
