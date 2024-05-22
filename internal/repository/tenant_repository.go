package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) port.ITenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) CreateTenant(tenant *domain.Tenant) (*domain.Tenant, error) {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	if err := r.db.Create(&tenantModel).Error; err != nil {
		return nil, err
	}

	tenant.Id = tenantModel.Id
	return tenant, nil
}

func (r *TenantRepository) FindTenantByName(name string) (*domain.Tenant, error) {
	return nil, nil
}
