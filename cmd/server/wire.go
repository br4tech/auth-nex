package main

import (
	handler "github.com/br4tech/auth-nex/internal/infra/delivery"
	"github.com/br4tech/auth-nex/internal/infra/repository"
	usecase "github.com/br4tech/auth-nex/internal/usecase/auth"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeAuthHandler(db *gorm.DB) *handler.AuthHandler {
	wire.Build(
		repository.NewAuthRepository,
		usecase.NewAuthUseCase,
		handler.NewAuthHandler,
	)

	return &handler.AuthHandler{}
}
