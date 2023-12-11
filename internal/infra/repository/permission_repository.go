package repository

import (
	"github.com/br4tech/auth-nex/internal/infra/model"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) IPermissionRepository {
	return &PermissionRepository{db: db}
}

func FindRoleByName(name string) (*model.User, error) {
	return nil, nil
}

func CreateRole(role *model.Role) error {
	return nil
}
