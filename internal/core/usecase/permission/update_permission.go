package permission

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type UpdatePermissionUseCase struct {
	permissionRepo port.IPermissionRepository
}

func NewUpdatePermissionUseCase(permissionRepo port.IPermissionRepository) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{
		permissionRepo: permissionRepo,
	}
}

func (uc *UpdatePermissionUseCase) Execute(permission *domain.Permission) error {
	permissionModel := new(model.Permission)
	permissionModel.FromDomain(permission)

	return uc.permissionRepo.Update(permissionModel)
}
