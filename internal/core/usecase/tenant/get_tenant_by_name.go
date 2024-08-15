package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetTenantByNameUseCase struct {
	tenantRepo port.ITenantRepository
}

func NewGetTenantByNameUseCase(tenantRepo port.ITenantRepository) *GetTenantByNameUseCase {
	return &GetTenantByNameUseCase{
		tenantRepo: tenantRepo,
	}
}

func (uc *GetTenantByNameUseCase) Execute(name string) (*domain.Tenant, error) {
	return uc.tenantRepo.FindByName(name)
}
