//go:build wireinject
// +build wireinject

package main

import (
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/core/usecase/auth"
	"github.com/br4tech/auth-nex/internal/core/usecase/company"
	"github.com/br4tech/auth-nex/internal/core/usecase/tenant"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/repository"
	"github.com/google/wire"
)

func InitializeUserHandler(db port.IDatabase) port.IUserHandler {
	wire.Build(
		repository.NewUserRepository,
		auth.NewAuthUseCase,
		handler.NewUserHandler,
		wire.Bind(new(port.IUserHandler), new(*handler.UserHandler)),
	)

	return &handler.UserHandler{}
}

func InitializeTenantHandler(db port.IDatabase) port.ITenantHandler {
	wire.Build(
		repository.NewTenantRepository,
		repository.NewCompanyRepository,
		repository.NewUserRepository,
		company.NewCompanyUseCase,
		auth.NewAuthUseCase,
		tenant.NewTenantUseCase,
		handler.NewTenantHandler,

		wire.Bind(new(port.ITenantHandler), new(*handler.TenantHandler)),
	)

	return &handler.TenantHandler{}
}
