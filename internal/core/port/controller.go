package port

import "github.com/labstack/echo/v4"

type (
	IUserHandler interface {
		CreateUser(c echo.Context) error
		GenerateToken(c echo.Context) error
		ValidateAccessToken(c echo.Context) error
	}

	ITenantHandler interface {
		CreateTenant(c echo.Context) error
	}
)
