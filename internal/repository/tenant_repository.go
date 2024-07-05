package repository

import (
	"context"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type TenantRepository struct {
	db port.IDatabase[model.Tenant]
}

func NewTenantRepository(db port.IDatabase[model.Tenant]) port.ITenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) CreateTenant(tenant *domain.Tenant) (*domain.Tenant, error) {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	_, err := r.db.Create(context.Background(), tenantModel)
	if err != nil {
		return nil, err
	}
	tenant.Id = tenantModel.Id
	return tenant, nil
}

func (r *TenantRepository) FindTenantByName(name string) (*domain.Tenant, error) {
	tenant, err := r.db.FindOne(context.Background(), "name=?", name)
	if err != nil {
		return nil, err
	}

	return tenant.ToDomain(), nil
}
