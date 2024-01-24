package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type (
	IPermissionRepository interface {
		CreateRole(role *domain.Role) error
		FindRoleByName(name string) (*domain.User, error)
	}

	IUserRepository interface {
		FindUserByEmail(email string) (*domain.User, error)
		CreateUser(user *domain.User) error
	}

	ITenantRepository interface {
		CreateTenant(tenant *domain.Tenant) (*domain.Tenant, error)
		FindTenantByName(name string) (*domain.Tenant, error)
	}
)
