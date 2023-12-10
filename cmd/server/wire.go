package main

import (
	handler "github.com/br4tech/auth-nex/internal/auth/delivery"
	"github.com/br4tech/auth-nex/internal/auth/repository"
	"github.com/br4tech/auth-nex/internal/auth/usecase"
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
