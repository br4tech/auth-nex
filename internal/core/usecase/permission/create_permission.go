package permission

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreatePermissionUseCase struct {
	permissionRepo port.IPermissionRepository
}

func NewCreatePermissionUseCase(permissionRepo port.IPermissionRepository) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		permissionRepo: permissionRepo,
	}
}

func (uc *CreatePermissionUseCase) Execute(permission *domain.Permission) (*domain.Permission, error) {
	permission, err := uc.permissionRepo.Create(permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}
