package repositories

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type TenantRepository struct {
	tenantAdapter port.IPostgreDatabase[model.Tenant]
}

func NewTenantRepository(tenantAdapter port.IPostgreDatabase[model.Tenant]) port.ITenantRepository {
	return &TenantRepository{
		tenantAdapter: tenantAdapter,
	}
}

func (r *TenantRepository) Create(tenant *domain.Tenant) error {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	tenantId, err := r.tenantAdapter.Create(*tenantModel)
	if err != nil {
		return err
	}

	tenant.Id = tenantId

	return nil
}

func (r *TenantRepository) FindByName(name string) (*domain.Tenant, error) {
	tenantEntity, err := r.tenantAdapter.FindBy("name=?", name)
	if err != nil {
		return nil, err
	}

	tenantModel, ok := any(tenantEntity).(*model.Tenant)
	if !ok {
		return nil, errors.New("failed to convert entity to tenant model")
	}

	tenantDomain := tenantModel.ToDomain()

	return tenantDomain, nil
}
