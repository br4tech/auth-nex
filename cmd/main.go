package main

import (
	"database/sql"
	"fmt"

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
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("Erro de conexao com o banco")
	}
	db.SetMaxOpenConns(10)

	tenantAdapter := adapter.NewPostgresAdapter[model.Tenant](db)
	tenantRepository := repositories.NewTenantRepository(tenantAdapter)

	// profileRepository := repositories.NewProfileRepository(postgresAdapter.(port.IDatabase[model.Profile]))
	userRepository := repositories.NewUserRepository(postgresAdapter.(port.IPostgreDatabase[model.User]))
	companyRepository := repositories.NewCompanyRepository(postgresAdapter.(port.IPostgreDatabase[model.Company]))

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
