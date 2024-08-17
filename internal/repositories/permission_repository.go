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

func (repo *permissionRepositoryImpl) Create(permission *domain.Permission) error {
	permissionModel := new(model.Permission)
	permissionModel.FromDomain(permission)

	return repo.db.Create(permissionModel).Error
}

func (repo *permissionRepositoryImpl) FindById(id int)(*domain.Permission, error){
	var permission model.Permission
	result := repo.db.First(&permission, id)

	return permission.ToDomain(), result.Error
}

func (repo *permissionRepositoryImpl) Update(permission *domain.Permission) error {
	permissionModel := new(model.Permission)
	permissionModel.FromDomain(permission)
	
	return repo.db.Save(permissionModel).Error
}

func (repo *permissionRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Permission{}, id).Error
}