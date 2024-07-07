package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type PermissionRepository struct {
	db port.IDatabase
}

func NewPermissionRepository(db port.IDatabase) port.IPermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindRoleByName(name string) (*domain.Profile, error) {
	var roleModel model.Profile

	if err := r.db.Where("name=?", name).First(&roleModel).Error; err != nil {
		return nil, err
	}

	return roleModel.ToDomain(), nil
}

func (r *PermissionRepository) CreateRole(role *domain.Profile) (*domain.Profile, error) {
	roleModel := new(model.Profile)
	roleModel.FromDomain(role)

	_, err := r.db.Create(roleModel)
	if err != nil {
		return nil, err
	}

	return roleModel.ToDomain(), nil
}
