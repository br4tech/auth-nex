package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type permissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) port.IPermissionRepository {
	return &permissionRepositoryImpl{
		db: db,
	}
}

func (repo *permissionRepositoryImpl) Create(permission *domain.Permission) (*domain.Permission, error) {
	permissionModel := new(model.Permission)
	permissionModel.FromDomain(permission)

	if err := repo.db.Create(permissionModel).Error; err != nil {
		return nil, err
	}

	return permissionModel.ToDomain(), nil
}

func (repo *permissionRepositoryImpl) FindById(id int) (*domain.Permission, error) {
	var permission model.Permission
	result := repo.db.First(&permission, id)

	return permission.ToDomain(), result.Error
}

func (repo *permissionRepositoryImpl) Update(permission *domain.Permission) (*domain.Permission, error) {
	permissionModel := new(model.Permission)
	permissionModel.FromDomain(permission)

	if err := repo.db.Save(permissionModel).Error; err != nil {
		return nil, err
	}

	return permissionModel.ToDomain(), nil
}

func (repo *permissionRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Permission{}, id).Error
}
