package permission

import "github.com/br4tech/auth-nex/internal/core/port"

type DeletePermissionUseCase struct {
	permissionRepo port.IPermissionRepository
}

func NewDeletePermissionUseCase(permissionRepo port.IPermissionRepository) *DeletePermissionUseCase {
	return &DeletePermissionUseCase{
		permissionRepo: permissionRepo,
	}
}

func (uc *DeletePermissionUseCase) Execute(id int) error {
	return uc.permissionRepo.Delete(id)
}
