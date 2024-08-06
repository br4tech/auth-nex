package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type (
	IProfileRepository interface {
		FindByName(name string) (*domain.Profile, error)
		Create(role *domain.Profile) (*domain.Profile, error)
	}

	IUserRepository interface {
		FindUserByEmail(email string) (*domain.User, error)
		FindByPhone(phone string) (*domain.User, error)
		CreateUser(user *domain.User) (*domain.User, error)
	}

	ICompanyRepository interface {
		FindById(id int) (*domain.Company, error)
		Create(company *domain.Company) (*domain.Company, error)
	}

	ITenantRepository interface {
		Create(tenant *domain.Tenant) (*domain.Tenant, error)
		FindByName(name string) (*domain.Tenant, error)
	}
)
