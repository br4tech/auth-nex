package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetTenantByIdUseCase struct {
	tenantRepo port.ITenantRepository
}

func NewGetTenantByIdUseCase(tenantRepo port.ITenantRepository) port.IGetTenantByIdUseCase {
	return &GetTenantByIdUseCase{
		tenantRepo: tenantRepo,
	}
}

func (uc *GetTenantByIdUseCase) Execute(id int) (*domain.Tenant, error) {
	tenant, err := uc.tenantRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}
