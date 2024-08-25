package port

import "github.com/labstack/echo/v4"

type (
	ITenantHandler interface {
		Create(c echo.Context) error
	}
)
