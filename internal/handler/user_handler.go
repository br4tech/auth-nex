package handler

import (
	"net/http"

	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase port.IUserUseCase
}

func NewUserHandler(userUseCase port.IUserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	reqBody := new(dto.UserDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := h.userUseCase.CreateUser(reqBody.ToDomain())
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create user")
	}

	return HandlerResponse(c, http.StatusOK, "Created user with successfully!!")
}

func (h *UserHandler) GenerateToken(c echo.Context) error {
	reqBody := new(dto.UserTokenDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	token, err := h.userUseCase.Authenticate(reqBody)

	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to authenticate")
	}

	return HandlerResponse(c, http.StatusOK, *token)
}

func (h *UserHandler) ValidateAccessToken(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	claims, err := h.userUseCase.ValidateAccessToken(token)
	if err != nil {
		return HandlerResponse(c, http.StatusUnauthorized, "Token inválido")
	}

	_ = claims

	return HandlerResponse(c, http.StatusOK, "token valido")
}
