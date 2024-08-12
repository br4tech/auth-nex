package permission

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type PermissionUseCase struct {
	permissionRepository port.IProfileRepository
}

func NewPermissionUseCase(permissionRepository port.IProfileRepository) port.IPermissionUseCase {
	return &PermissionUseCase{
		permissionRepository: permissionRepository,
	}
}

func (uc *PermissionUseCase) FindByName(name string) (*domain.Profile, error) {
	profile, err := uc.permissionRepository.FindByName(name)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (uc *PermissionUseCase) Create(name string) (*domain.Profile, error) {
	profileModel := &model.Profile{Name: name}

	roleCreated, err := uc.permissionRepository.Create(profileModel.ToDomain())

	if err != nil {
		return nil, err
	}

	return roleCreated, nil
}
