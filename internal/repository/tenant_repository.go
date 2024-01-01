package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/jinzhu/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) port.ITenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) CreateTenant(model *model.Tenant) (*domain.Tenant, error) {
	return nil, nil
}
func (r *TenantRepository) FindTenantByName(name string) (*domain.Tenant, error) {
	return nil, nil
}
