package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type IPermissionRepository interface {
	FindRoleByName(name string) (*domain.Role, error)
	CreateRole(role *domain.Role) (*domain.Role, error)
}

type IUserRepository interface {
	FindUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
}

type ICompanyRepository interface {
	FindCompanyById(id int) (*domain.Company, error)
	CreateCompany(company *domain.Company) (*domain.Company, error)
}

type ITenantRepository interface {
	CreateTenant(tenant *domain.Tenant) (*domain.Tenant, error)
	FindTenantByName(name string) (*domain.Tenant, error)
}
