package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type companyRepositoryImpl struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) port.ICompanyRepository {
	return &companyRepositoryImpl{
		db: db,
	}
}

func (repo *companyRepositoryImpl) Create(company *model.Company) error {
	return repo.db.Create(company).Error
}

func (repo *companyRepositoryImpl) FindById(id int) (*model.Company, error) {
	var company model.Company
	result := repo.db.First(&company, id)

	return &company, result.Error
}

func (repo *companyRepositoryImpl) Update(company *model.Company) error {
	return repo.db.Save(company).Error
}

func (repo *companyRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Company{}, id).Error
}
