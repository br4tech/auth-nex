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

func (uc *PermissionUseCase) FindProfileByName(name string) (*domain.Profile, error) {
	profile, err := uc.permissionRepository.FindProfileByName(name)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (uc *PermissionUseCase) CreateProfile(name string) (*domain.Profile, error) {
	profileModel := &model.Profile{Name: name}

	roleCreated, err := uc.permissionRepository.CreateProfile(profileModel.ToDomain())

	if err != nil {
		return nil, err
	}

	return roleCreated, nil
}
