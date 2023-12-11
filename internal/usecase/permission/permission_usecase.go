package usecase

import (
	"github.com/br4tech/auth-nex/internal/infra/model"
	"github.com/br4tech/auth-nex/internal/infra/repository"
)

type PermissionUseCase struct {
	permissionRepository repository.IPermissionRepository
}

func NewPermissionUseCase(permissionRepository repository.IPermissionRepository) IPermissionUseCase {
	return &PermissionUseCase{
		permissionRepository: permissionRepository,
	}
}

func (uc *PermissionUseCase) CreateRole(name string) error {
	role := &model.Role{Name: name}
	return uc.permissionRepository.CreateRole(role)
}
