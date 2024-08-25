package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateTenantUseCase struct {
	tenantRepo port.ITenantRepository
}

func NewUpdateTenantUseCase(tenantRepo port.ITenantRepository) port.IUpdateTenantUseCase {
	return &UpdateTenantUseCase{
		tenantRepo: tenantRepo,
	}
}

func (uc *UpdateTenantUseCase) Execute(tenant *domain.Tenant) (*domain.Tenant, error) {
	tenant, err := uc.tenantRepo.Update(tenant)
	if err != nil {
		return nil, err
	}
	return tenant, err
}
