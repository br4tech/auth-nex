package port

import (
	"github.com/br4tech/auth-nex/internal/model"
)

type (
	ICompanyRepository interface {
		Create(company *model.Company) error
		FindById(id int) (*model.Company, error)
		Update(company *model.Company) error
		Delete(id int) error
	}

	IPermissionRepository interface {
		Create(permission *model.Permission) error
		FindById(id int) (*model.Permission, error)
		Update(permission *model.Permission) error
		Delete(id int) error
	}

	IProfileRepository interface {
		Create(profile *model.Profile) error
		FindById(id int) (*model.Profile, error)
		Upate(profile *model.Profile) error
		Delete(id int) error
	}

	ITenantRepository interface {
		Create(tenant *model.Tenant) error
		FindById(id int) (*model.Tenant, error)
		FindByName(name string) (*model.Tenant, error)
		Update(tenant *model.Tenant) error
		Delete(id int) error
	}

	IUserRepository interface {
		Create(user *model.User) error
		FindById(id int) (*model.User, error)
		Update(user *model.User) error
		Delete(id int) error
	}
)
