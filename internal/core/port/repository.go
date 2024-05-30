package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type (
	IPermissionRepository interface {
		FindRoleByName(name string) (*domain.Role, error)
		CreateRole(role *domain.Role) (*domain.Role, error)
	}

	IUserRepository interface {
		FindUserByEmail(email string) (*domain.User, error)
		CreateUser(user *domain.User) (*domain.User, error)
	}

	ICompanyRepository interface {
		FindCompanyById(id int) (*domain.Company, error)
		CreateCompany(company *domain.Company) (*domain.Company, error)
	}

	ITenantRepository interface {
		CreateTenant(tenant *domain.Tenant) (*domain.Tenant, error)
		FindTenantByName(name string) (*domain.Tenant, error)
	}
)
