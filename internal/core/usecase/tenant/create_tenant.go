package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
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
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	return uc.tenantRepository.Create(tenantModel)
}
