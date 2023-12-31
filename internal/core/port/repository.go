package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/model"
)

type (
	IPermissionRepository interface {
		CreateRole(role *domain.Role) error
		FindRoleByName(name string) (*domain.User, error)
	}

	IUserRepository interface {
		Authenticate(name string, tenantID uint) (*domain.User, error)
		CreateUser(userModel *model.User) error
	}

	ITenantRepository interface {
		CreateTenant(model *model.Tenant) (*domain.Tenant, error)
		FindTenantByName(name string) (*domain.Tenant, error)
	}
)
