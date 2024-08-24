package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewProfileFactory(name string) *domain.Profile {
	return &domain.Profile{
		Id:       1,
		Name:     name,
		TenantId: 1,
	}
}
