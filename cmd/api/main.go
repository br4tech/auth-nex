package main

import (
	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/adapter"
	"github.com/br4tech/auth-nex/internal/handler"
	"github.com/br4tech/auth-nex/server"
)

func main() {
	cfg := config.GetConfig()
	db := adapter.NewPostgresDatabase(&cfg)
	userHandler := handler.NewUserHandler()
	server.NewEchoServer(&cfg, db.GetDb(), userHandler).Start()
}
