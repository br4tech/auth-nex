package handler

import (
	"net/http"

	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authenticateUserUseCase port.IAuthenticateUserUseCase
}

func NewAuthHandler(authenticateUserUseCase port.IAuthenticateUserUseCase) *AuthHandler {
	return &AuthHandler{
		authenticateUserUseCase: authenticateUserUseCase,
	}
}

func (h *AuthHandler) Token(c echo.Context) error {
	reqBody := new(dto.UserTokenDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	token, err := h.authenticateUserUseCase.Execute(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to authenticate")
	}

	return HandlerResponse(c, http.StatusOK, *token)
}
