package port

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
)

type IUserUseCase interface {
	Authenticate(userReq *dto.UserTokenDTO) (*string, error)
	CreateUser(user *domain.User) (*domain.User, error)
	ValidateAccessToken(tokenString string) (*model.Claims, error)
}

type IPermissionUseCase interface {
	FindProfileByName(name string) (*domain.Profile, error)
	CreateProfile(name string) (*domain.Profile, error)
}

type ICompanyUseCase interface {
	FindCompanyById(id int) (*domain.Company, error)
	CreateCompany(company *domain.Company) (*domain.Company, error)
}

type ITenantUseCase interface {
	CreateTenantWithCompanyAndAdmin(tenant *dto.TenantDTO) (*domain.Tenant, error)
}
