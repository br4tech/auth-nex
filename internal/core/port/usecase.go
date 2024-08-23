package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
)

type IUserUseCase interface {
	Authenticate(userReq *dto.UserTokenDTO) (*string, error)
	Create(user *domain.User) (*domain.User, error)
	ValidateAccessToken(tokenString string) (*model.Claims, error)
}

type IPermissionUseCase interface {
	FindByName(name string) (*domain.Profile, error)
	Create(name string) (*domain.Profile, error)
}

type ICompanyUseCase interface {
	FindById(id int) (*domain.Company, error)
	Create(company *domain.Company) (*domain.Company, error)
}

type ITenantUseCase interface {
	CreateTenantWithCompanyAndAdmin(tenant *dto.TenantDTO) (*domain.Tenant, error)
}
