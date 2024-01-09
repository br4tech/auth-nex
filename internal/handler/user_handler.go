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

func NewUserHandler(userUseCase port.IUserUseCase) port.IUserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	reqBody := new(dto.UserDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := h.userUseCase.CreateUser(reqBody)
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create user")
	}

	return HandlerResponse(c, http.StatusOK, "Created user with successfuly!!")
}
