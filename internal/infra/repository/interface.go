package repository

import "github.com/br4tech/auth-nex/internal/infra/model"

type (
	IAuthRepository interface {
		Authenticate(username string, tenantID uint) (*model.User, error)
		Register(userModel *model.User) error
	}

	IPermissionRepository interface {
		FindRoleByName(name string) (*model.User, error)
		CreateRole(role *model.Role) error
	}
)
