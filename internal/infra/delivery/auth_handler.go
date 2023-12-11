package handler

import (
	"net/http"

	usecase "github.com/br4tech/auth-nex/internal/usecase/auth"
)

// AuthHandler manipula solicitações HTTP relacionadas à autenticação.
type AuthHandler struct {
	authUseCase usecase.IAuthUseCase
}

func NewAuthHandler(authUseCase usecase.IAuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

// AuthenticateHandler autentica um usuário.
func (h *AuthHandler) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	// Implemente a lógica de manipulação da solicitação de autenticação.
}

// RegisterHandler registra um novo usuário.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Implemente a lógica de manipulação da solicitação de registro.
}
