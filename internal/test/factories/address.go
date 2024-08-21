package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewAddressFactory() *domain.Address {
	return &domain.Address{
		Street:     "Rua de Teste",
		Number:     "123",
		Complement: "Apto 42",
		District:   "Centro",
		City:       "São Paulo",
		State:      "SP",
		ZipCode:    "01001-000",
	}
}
