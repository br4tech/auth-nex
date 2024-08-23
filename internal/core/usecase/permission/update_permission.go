package permission

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdatePermissionUseCase struct {
	permissionRepo port.IPermissionRepository
}

func NewUpdatePermissionUseCase(permissionRepo port.IPermissionRepository) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		permissionRepo: permissionRepo,
	}
}

func (uc *UpdatePermissionUseCase) Execute(permission *domain.Permission) (*domain.Permission, error) {
	permission, err := uc.permissionRepo.Update(permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
