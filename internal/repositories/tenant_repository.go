package repositories

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type TenantRepository[T port.IModel] struct {
	db port.IDatabase[T]
}

func NewTenantRepository[T port.IModel](db port.IDatabase[T]) port.ITenantRepository {
	return &TenantRepository[T]{db: db}
}

func (r *TenantRepository[T]) Create(tenant *domain.Tenant) (*domain.Tenant, error) {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	tenantEntity, ok := any(tenantModel).(T)
	if !ok {
		return nil, errors.New("entity type invalid to tenant")
	}

	tenantId, err := r.db.Create(tenantEntity)
	if err != nil {
		return nil, err
	}

	tenant.Id = tenantId

	return tenant, nil
}

func (r *TenantRepository[T]) FindByName(name string) (*domain.Tenant, error) {
	tenantEntity, err := r.db.FindBy("name", name)
	if err != nil {
		return nil, err
	}

	tenantModel, ok := any(tenantEntity[0]).(*model.Tenant)
	if !ok {
		return nil, errors.New("failed to convert entity to tenant model")
	}

	tenantDomain := tenantModel.ToDomain()

	return tenantDomain, nil
}
