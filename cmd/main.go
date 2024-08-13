package main

import (
	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/adapter"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/core/usecase/company"
	"github.com/br4tech/auth-nex/internal/core/usecase/tenant"
	"github.com/br4tech/auth-nex/internal/core/usecase/user"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/br4tech/auth-nex/internal/repositories"
	"github.com/br4tech/auth-nex/server"
)

func main() {
	cfg := config.GetConfig()
	postgresAdapter := adapter.NewPostgresAdapter[port.IModel](&cfg)

	companyRepository := repositories.NewCompanyRepository(postgresAdapter.(port.IDatabase[model.Company]))
	// profileRepository := repositories.NewProfileRepository(postgresAdapter.(port.IDatabase[model.Profile]))
	tenantRepository := repositories.NewTenantRepository(postgresAdapter.(port.IDatabase[model.Tenant]))
	userRepository := repositories.NewUserRepository(postgresAdapter.(port.IDatabase[model.User]))

	userUseCase := user.NewAuthUseCase(userRepository)
	companyUseCase := company.NewCompanyUseCase(companyRepository)
	tenantUseCase := tenant.NewTenantUseCase(tenantRepository, companyUseCase, userUseCase)

	userHandler := handler.NewUserHandler(userUseCase)
	tenantHandler := handler.NewTenantHandler(tenantUseCase)

	server.NewEchoServer(
		&cfg,
		postgresAdapter.GetDb(),
		userHandler,
		tenantHandler,
	).Start()

}
