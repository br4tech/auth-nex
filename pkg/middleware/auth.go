package middleware

import (
	"net/http"

	"github.com/br4tech/auth-nex/internal/auth/usecase"
)

// AuthMiddleware é um middleware para autenticação.
func AuthMiddleware(next http.Handler, authUseCase usecase.AuthUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implemente a lógica de autenticação aqui.
		// Se o usuário estiver autenticado, chame next.ServeHTTP(w, r)
		// Caso contrário, retorne um erro ou redirecione conforme necessário.
	})
}
