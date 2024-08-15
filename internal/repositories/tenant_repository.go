package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type tenantRepositoryImp struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) port.ITenantRepository {
	return &tenantRepositoryImp{
		db: db,
	}
}

func (repo *tenantRepositoryImp) Create(tenant *model.Tenant) error {
	return repo.db.Create(tenant).Error
}

func (repo *tenantRepositoryImp) FindById(id int) (*model.Tenant, error) {
	var tenant model.Tenant
	result := repo.db.First(&tenant, id)

	return &tenant, result.Error
}

func (repo *tenantRepositoryImp) Update(tenant *model.Tenant) error {
	return repo.db.Save(tenant).Error
}

func (repo *tenantRepositoryImp) Delete(id int) error {
	return repo.db.Delete(&model.Tenant{}, id).Error
}
