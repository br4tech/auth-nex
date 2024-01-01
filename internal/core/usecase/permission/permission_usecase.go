package usecase

import (
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type PermissionUseCase struct {
	permissionRepository port.IPermissionRepository
}

func NewPermissionUseCase(permissionRepository port.IPermissionRepository) port.IPermissionUseCase {
	return &PermissionUseCase{
		permissionRepository: permissionRepository,
	}
}

func (uc *PermissionUseCase) CreateRole(name string) error {
	role := &model.Role{Name: name}
	return uc.permissionRepository.CreateRole(role)
}
