package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
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

func (repo *tenantRepositoryImp) Create(tenant *domain.Tenant) error {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	return repo.db.Create(tenantModel).Error
}

func (repo *tenantRepositoryImp) FindById(id int) (*domain.Tenant, error) {
	var tenantModel model.Tenant
	result := repo.db.First(&tenantModel, id)

	return tenantModel.ToDomain(), result.Error
}

func (repo *tenantRepositoryImp) FindByName(name string) (*domain.Tenant, error) {
	var tenantModel model.Tenant
	result := repo.db.First(tenantModel, name)

	return tenantModel.ToDomain(), result.Error
}

func (repo *tenantRepositoryImp) Update(tenant *domain.Tenant) error {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	return repo.db.Save(tenant).Error
}

func (repo *tenantRepositoryImp) Delete(id int) error {
	return repo.db.Delete(&model.Tenant{}, id).Error
}
