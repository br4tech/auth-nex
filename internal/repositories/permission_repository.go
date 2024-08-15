``package repositories

import "gorm.io/gorm"

type permissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) port.IPermissionRepository{
	return &permissionRepositoryImpl{
		db: db
	}
}

func (repo *permissionRepositoryImpl) Create(permission *model.Permission) error {
	return repo.db.Create(permission).Error
}

func (repo *permissionRepositoryImpl) FindById(id int)(*model.Permission, error){
	var permission model.Permission
	result := repo.db.First(&permission, id)

	return &permission, result.Error
}

func (repo *permissionRepositoryImpl) Update(permission *model.Permission) error {
	return repo.db.Save(permission).Error
}

func (repo *permissionRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Permission{}, id).Error
}