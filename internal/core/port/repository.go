package port

import "github.com/br4tech/auth-nex/internal/core/domain"

type (
	ICompanyRepository interface {
		Create(company *domain.Company) (*domain.Company, error)
		FindById(id int) (*domain.Company, error)
		Update(company *domain.Company) (*domain.Company, error)
		Delete(id int) error
	}

	IPermissionRepository interface {
		Create(permission *domain.Permission) (*domain.Permission, error)
		FindById(id int) (*domain.Permission, error)
		Update(permission *domain.Permission) (*domain.Permission, error)
		Delete(id int) error
	}

	IProfileRepository interface {
		Create(profile *domain.Profile) (*domain.Profile, error)
		FindById(id int) (*domain.Profile, error)
		Upate(profile *domain.Profile) (*domain.Profile, error)
		Delete(id int) error
	}

	ITenantRepository interface {
		Create(tenant *domain.Tenant) (*domain.Tenant, error)
		FindById(id int) (*domain.Tenant, error)
		FindByName(name string) (*domain.Tenant, error)
		Update(tenant *domain.Tenant) (*domain.Tenant, error)
		Delete(id int) error
	}

	IUserRepository interface {
		Create(user *domain.User) (*domain.User, error)
		FindBy(filter map[string]interface{}) (*domain.User, error)
		FindById(id int) (*domain.User, error)
		FindByEmail(email string) (*domain.User, error)
		Update(user *domain.User) (*domain.User, error)
		Delete(id int) error
	}
)
