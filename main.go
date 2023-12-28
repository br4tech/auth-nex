package main

import (
	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/database"
	"github.com/br4tech/auth-nex/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
