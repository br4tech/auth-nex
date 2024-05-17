package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/br4tech/auth-nex/pkg/validator"

	"github.com/jinzhu/copier"
)

type TenantUseCase struct {
	tenantRepository port.ITenantRepository
}

func NewTenantUseCase(tenantRepository port.ITenantRepository) port.ITenantUseCase {
	return &TenantUseCase{tenantRepository: tenantRepository}
}

func (uc *TenantUseCase) CreateTenant(tenant *dto.TenantDTO) (*domain.Tenant, error) {
	tenantModel := &model.Tenant{}
	err := copier.Copy(tenantModel, tenant)

	if err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(tenantModel); err != nil {
		return nil, err
	}

	uc.tenantRepository.CreateTenant(tenantModel.ToDomain())

	return tenantModel.ToDomain(), nil
}
