package permission

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
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

func (uc *PermissionUseCase) FindRoleByName(name string) (*domain.Role, error) {
	role, err := uc.permissionRepository.FindRoleByName(name)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (uc *PermissionUseCase) CreateRole(name string) (*domain.Role, error) {
	roleModel := &model.Role{Name: name}

	roleCreated, err := uc.permissionRepository.CreateRole(roleModel.ToDomain())

	if err != nil {
		return nil, err
	}

	return roleCreated, nil
}
