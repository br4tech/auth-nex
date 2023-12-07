package handler

import (
	"net/http"

	"github.com/br4tech/auth-nex/internal/auth/usecase"
)

// AuthHandler manipula solicitações HTTP relacionadas à autenticação.
type AuthHandler struct {
	AuthUseCase usecase.AuthUseCase
}

// AuthenticateHandler autentica um usuário.
func (h *AuthHandler) AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	// Implemente a lógica de manipulação da solicitação de autenticação.
}

// RegisterHandler registra um novo usuário.
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Implemente a lógica de manipulação da solicitação de registro.
}
