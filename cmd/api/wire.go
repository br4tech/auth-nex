package main

import (
	usecase "github.com/br4tech/auth-nex/internal/core/usecase/auth"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/repository"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserHandler(db *gorm.DB) *handler.UserHandler {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewAuthUseCase,
		handler.NewUserHandler,
	)

	return &handler.UserHandler{}
}
