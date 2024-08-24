package main

import (
	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/adapter"
	"github.com/br4tech/auth-nex/internal/core/usecase/tenant"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/repositories"
	"github.com/br4tech/auth-nex/server"
)

func main() {
	cfg := config.GetConfig()

	db, err := adapter.NewPostgresAdapter(cfg)
	if err != nil {
		panic("Erro de conexao com o banco")
	}

	tenantRepository := repositories.NewTenantRepository(db)
	companyRepository := repositories.NewCompanyRepository(db)
	userRepository := repositories.NewUserRepository(db)
	profileRepository := repositories.NewProfileRepository(db)

	createTenantUseCase := tenant.NewCreateTenantUseCase(tenantRepository, companyRepository, userRepository, profileRepository)
	updateTenantUseCase := tenant.NewUpdateTenantUseCase(tenantRepository)
	getTenantByIdUseCase := tenant.NewGetTenantByIdUseCase(tenantRepository)
	getTenantByNameUseCase := tenant.NewGetTenantByNameUseCase(tenantRepository)

	tenantHandler := handler.NewTenantHandler(createTenantUseCase, updateTenantUseCase, getTenantByIdUseCase, getTenantByNameUseCase)

	server.NewEchoServer(
		&cfg,
		db,
		tenantHandler,
	).Start()

}
