package repositories

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
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

func (repo *companyRepositoryImpl) Create(company *domain.Company) error {
	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	return repo.db.Create(companyModel).Error
}

func (repo *companyRepositoryImpl) FindById(id int) (*domain.Company, error) {
	var company model.Company
	result := repo.db.First(&company, id)

	return company.ToDomain(), result.Error
}

func (repo *companyRepositoryImpl) Update(company *domain.Company) error {
	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	return repo.db.Save(companyModel).Error
}

func (repo *companyRepositoryImpl) Delete(id int) error {
	return repo.db.Delete(&model.Company{}, id).Error
}
