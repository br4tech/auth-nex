package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewUserFactory(role string) *domain.User {
	return &domain.User{
		Id:        1,
		Name:      "Usuário Teste",
		Email:     "usuario@teste.com",
		CPF:       "12345678901",
		Password:  "senha123",
		Phone:     "+55 11 987654321",
		TenantId:  1,
		Role:      role,
		ProfileId: 1,
	}
}
