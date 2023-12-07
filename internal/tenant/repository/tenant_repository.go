package repository

import "github.com/br4tech/auth-nex/internal/tenant/model"

type TenantRepository interface {
	FindTenantByName(name string) (*model.Tenant, error)
	CreateTenant(tenant *model.Tenant) error
}
