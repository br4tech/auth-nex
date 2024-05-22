package server

import (
	"fmt"

	"github.com/br4tech/auth-nex/config"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app           *echo.Echo
	db            *gorm.DB
	cfg           *config.Config
	userHandler   port.IUserHandler
	tenantHandler port.ITenantHandler
}

func NewEchoServer(
	cfg *config.Config,
	db *gorm.DB,
	userHandler port.IUserHandler,
	tenantHandler port.ITenantHandler,
) Server {
	return &echoServer{
		app:           echo.New(),
		db:            db,
		cfg:           cfg,
		userHandler:   userHandler,
		tenantHandler: tenantHandler,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	s.app.POST("/tenant", s.tenantHandler.CreateTenant)

	s.app.POST("/token", s.userHandler.GenerateToken)
	s.app.POST("/user", s.userHandler.CreateUser)
	s.app.GET("/validate_token", s.userHandler.ValidateAccessToken)

	serverUrl := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
