package usecase

import (
	"github.com/br4tech/auth-nex/internal/permission/model"
	"github.com/br4tech/auth-nex/internal/permission/repository"
)

type PermissionUseCaseImpl struct {
	permissionRepo repository.PermissionRepository
}

func NewPermissionUseCase(permissionRepo repository.PermissionRepository) PermissionUseCase {
	return &PermissionUseCaseImpl{
		permissionRepo: permissionRepo,
	}
}

func (uc *PermissionUseCaseImpl) CreateRole(name string) error {
	role := &model.Role{Name: name}
	return uc.permissionRepo.CreateRole(role)
}
