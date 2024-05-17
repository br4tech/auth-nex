package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
)

type (
	IUserUseCase interface {
		Authenticate(userReq *dto.UserTokenDTO) (*string, error)
		CreateUser(user *dto.UserDTO) (*domain.User, error)
		ValidateAccessToken(tokenString string) (*model.Claims, error)
	}

	IPermissionUseCase interface {
		CreateRole(name string) error
	}

	ICompanyUseCase interface {
		FindCompanyById(id int) (*domain.Company, error)
		CreateCompany(company *domain.Company) (*domain.Company, error)
	}

	ITenantUseCase interface {
		CreateTenant(tenant *dto.TenantDTO) (*domain.Tenant, error)
	}
)
