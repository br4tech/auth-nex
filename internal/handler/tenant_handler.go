package handler

import (
	"net/http"

	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/labstack/echo/v4"
)

type TenantHandler struct {
	tenantUseCase port.ITenantUseCase
}

func NewTenantHandler(tenantUseCase port.ITenantUseCase) *TenantHandler {
	return &TenantHandler{
		tenantUseCase: tenantUseCase,
	}
}

func (h *TenantHandler) CreateTenant(c echo.Context) error {
	reqBody := new(dto.TenantDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := h.tenantUseCase.CreateTenantWithCompanyAndAdmin(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed do create tenant")
	}

	return HandlerResponse(c, http.StatusOK, "Created tenant with sucessfully")
}
