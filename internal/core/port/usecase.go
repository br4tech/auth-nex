package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
)

type (
	ICreateTenantUseCase interface {
		Execute(tenant *domain.Tenant) error
	}

	IGetTenantByIdUseCase interface {
		Execute(id int) (*domain.Tenant, error)
	}

	IGetTenantByNameUseCase interface {
		Execute(name string) (*domain.Tenant, error)
	}

	IUpdateTenantUseCase interface {
		Execute(tenant *domain.Tenant) (*domain.Tenant, error)
	}

	ICreateUserUsecase interface {
		Execute(user *domain.User) error
	}

	IDeleteUsecase interface {
		Execute(id int) error
	}

	IGetUserByEmailUseCase interface {
		Execute(email string) (*domain.User, error)
	}

	IUpdateUserUsecase interface {
		Execute(user *domain.User) (*domain.User, error)
	}
)
