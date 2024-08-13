package main

import (
	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/adapter"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

func main() {
	cfg := config.GetConfig()
	postgresAdapter := adapter.NewPostgresAdapter[port.IModel](&cfg)

	AuthnexMigrate(postgresAdapter)
}

func AuthnexMigrate(db port.IDatabase[port.IModel]) {
	db.GetDb().Migrator().CreateTable(
		&model.Tenant{},
		&model.User{},
		&model.Company{},
		&model.Activity{},
		&model.Address{},
		&model.Partner{},
	)
}
