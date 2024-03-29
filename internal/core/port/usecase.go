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

	ITenantUseCase interface {
		CreateTenant(name string) error
	}
)
