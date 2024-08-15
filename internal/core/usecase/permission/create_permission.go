package permission

import "github.com/br4tech/auth-nex/internal/core/port"

type CreatePermissionUseCase struct {
	permissionRepo port.IProfileRepository
}

func NewCreatePermissionUseCase(permissionRepo port.IProfileRepository) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{
		permissionRepo: permissionRepo,
	}
}
