package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) port.IPermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) FindRoleByName(name string) (*domain.Role, error) {
	var roleModel model.Role

	if err := r.db.Where("name=?", name).First(&roleModel).Error; err != nil {
		return nil, err
	}

	return roleModel.ToDomain(), nil
}

func (r *PermissionRepository) CreateRole(role *domain.Role) (*domain.Role, error) {
	roleModel := new(model.Role)
	roleModel.FromDomain(role)

	if err := r.db.Create(&roleModel).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
