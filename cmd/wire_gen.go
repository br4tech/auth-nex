// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/br4tech/auth-nex/internal/core/usecase/auth"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/internal/repository"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeUserHandler(db *gorm.DB) *handler.UserHandler {
	iUserRepository := repository.NewUserRepository(db)
	iUserUseCase := usecase.NewAuthUseCase(iUserRepository)
	userHandler := handler.NewUserHandler(iUserUseCase)
	return userHandler
}
