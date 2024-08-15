package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateTenantUseCase struct {
	tenantRepo port.ITenantRepository
}

func NewUpdateTenantUseCase(tenantRepo port.ITenantRepository) *UpdateTenantUseCase {
	return &UpdateTenantUseCase{
		tenantRepo: tenantRepo,
	}
}

func (uc *UpdateTenantUseCase) Execute(tenant domain.Tenant) error {
	return nil
}
