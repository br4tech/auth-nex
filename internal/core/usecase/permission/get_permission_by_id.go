package permission

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetPermissionById struct {
	permissionRepo port.IPermissionRepository
}

func NewGetPermissionById(permissionRepo port.IPermissionRepository) *GetPermissionById {
	return &GetPermissionById{
		permissionRepo: permissionRepo,
	}
}

func (uc *GetPermissionById) Execute(id int) (*domain.Permission, error) {
	permission, err := uc.permissionRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return permission.ToDomain(), nil
}
