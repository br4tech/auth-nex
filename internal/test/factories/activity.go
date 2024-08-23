package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewActivityFactory() *domain.Activity {
	return &domain.Activity{
		CNAE:        "0111301",          // Ou outro código CNAE padrão
		Description: "Cultivo de arroz", // Ou outra descrição padrão
	}
}
