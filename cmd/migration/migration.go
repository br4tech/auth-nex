package main

import (
	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/adapter"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

func main() {
	cfg := config.GetConfig()
	db := adapter.NewPostgresDatabase(&cfg)

	AuthnexMigrate(db)
}

func AuthnexMigrate(db port.IDatabase) {
	db.GetDb().Migrator().CreateTable(
		&model.Tenant{},
		&model.User{},
		&model.Company{},
		&model.Activity{},
		&model.Address{},
		&model.Partner{},
	)
}
