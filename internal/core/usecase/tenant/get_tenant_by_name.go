package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetTenantByNameUseCase struct {
	tenantRepo port.ITenantRepository
}

func NewGetTenantByNameUseCase(tenantRepo port.ITenantRepository) port.IGetTenantByNameUseCase {
	return &GetTenantByNameUseCase{
		tenantRepo: tenantRepo,
	}
}

func (uc *GetTenantByNameUseCase) Execute(name string) (*domain.Tenant, error) {
	tenant, err := uc.tenantRepo.FindByName(name)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}
