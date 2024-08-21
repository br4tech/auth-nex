package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewPermissionFactory(name string) *domain.Permission {
	return &domain.Permission{
		Id:   1,
		Name: name,
	}
}
