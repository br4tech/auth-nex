package main

import (
	"database/sql"
	"fmt"

	"github.com/br4tech/auth-nex/config"
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

	// tenantAdapter := adapter.NewPostgresAdapter[model.Tenant](db)
	// tenantRepository := repositories.NewTenantRepository(tenantAdapter)

	// // profileRepository := repositories.NewProfileRepository(postgresAdapter.(port.IDatabase[model.Profile]))
	// userRepository := repositories.NewUserRepository(postgresAdapter.(port.IPostgreDatabase[model.User]))
	// companyRepository := repositories.NewCompanyRepository(postgresAdapter.(port.IPostgreDatabase[model.Company]))

	// userUseCase := user.NewAuthUseCase(userRepository)
	// companyUseCase := company.NewCompanyUseCase(companyRepository)
	// tenantUseCase := tenant.NewTenantUseCase(tenantRepository, companyUseCase, userUseCase)

	// userHandler := handler.NewUserHandler(userUseCase)
	// tenantHandler := handler.NewTenantHandler(tenantUseCase)

	// server.NewEchoServer(
	// 	&cfg,
	// 	postgresAdapter.GetDb(),
	// 	userHandler,
	// 	tenantHandler,
	// ).Start()

}
