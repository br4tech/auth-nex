package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/model"
)

type (
	IProfileRepository interface {
		Create(profile *model.Profile) error
		FindById(id int) (*model.Profile, error)
		Upate(profile *model.Profile) error
		Delete(id int) error
	}

	IUserRepository interface {
		Create(user *model.User) error
		FindById(id int) (*model.User, error)
		Update(user *model.User) error
		Delete(id int) error
	}

	ICompanyRepository interface {
		Create(company *model.Company) error
		FindById(id int) (*model.Company, error)
		Update(company *model.Company) error
		Delete(id int) error
	}

	ITenantRepository interface {
		Create(tenant *domain.Tenant) error
		FindByName(name string) (*domain.Tenant, error)
	}
)
