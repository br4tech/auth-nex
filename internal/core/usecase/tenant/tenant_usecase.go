package usecase

import "github.com/br4tech/auth-nex/internal/core/port"

type TenantUseCase struct {
	tenantRepository port.ITenantRepository
}

func NewTenantUseCase(tenantRepository port.ITenantRepository) port.ITenantUseCase {
	return &TenantUseCase{tenantRepository: tenantRepository}
}

func (uc *TenantUseCase) CreateTenant(name string) error {
	return nil
}
