package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateTenantUseCase struct {
	tenantRepository port.ITenantRepository
}

func NewCreateTenantUseCase(tenantRepository port.ITenantRepository) *CreateTenantUseCase {
	return &CreateTenantUseCase{
		tenantRepository: tenantRepository,
	}
}

func (uc *CreateTenantUseCase) Execute(tenant *domain.Tenant) error {
	return uc.tenantRepository.Create(tenant)
}
