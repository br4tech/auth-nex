//go:build wireinject
// +build wireinject

package main

import (
  "github.com/br4tech/auth-nex/internal/core/usecase/auth"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)


//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/br4tech/go-with-gemini/config"
	"github.com/br4tech/go-with-gemini/internal/adapter"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/core/usecase"
	"github.com/br4tech/go-with-gemini/internal/handler"
	"github.com/br4tech/go-with-gemini/internal/repository"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB) *handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		auth.NewAuthUseCase,
		handler.NewUserHandler,
	)

	return &handler.UserHandler{}
}

func InitializeTenantHandler(db *gorm.DB) *handler.TenantHandler {
	wire.Build(
		repository.NewTenantRepository,
		tenant.NewTenantUseCase,
		handler.NewTenantHandler,
	)

	return &handler.TenantHandler{}
}
