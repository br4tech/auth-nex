package factories

import "github.com/br4tech/auth-nex/internal/core/domain"

func NewTenantFactory() *domain.Tenant {
	return &domain.Tenant{
		Id:        1,
		Name:      "Tenant Teste",
		Companies: []domain.Company{*NewCompanyFactory()},
		Users:     []domain.User{*NewUserFactory("system_admin")},
	}
}
