package adapter

import (
	"fmt"
	"log"

	"github.com/br4tech/auth-nex/config"
	"gorm.io/driver/postgres" // Adapte para o driver do seu banco de dados
	"gorm.io/gorm"
)

func NewPostgresAdapter(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.DBName, cfg.Db.Port, cfg.Db.SSLMode, cfg.Db.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("Connected to the database!") // Ou use um logger mais robusto

	return db, nil
}
