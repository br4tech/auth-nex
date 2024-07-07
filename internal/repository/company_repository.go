package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type CompanyRepository struct {
	db port.IDatabase
}

func NewCompanyRepository(db port.IDatabase) port.ICompanyRepository {
	return &CompanyRepository{db: db}
}

func (r CompanyRepository) FindCompanyById(id int) (*domain.Company, error) {
	var companyModel model.Company

	if err := r.db.Where("id=?", id).First(&companyModel).Error; err != nil {
		return nil, err
	}

	return companyModel.ToDomain(), nil
}

func (r CompanyRepository) CreateCompany(company *domain.Company) (*domain.Company, error) {
	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	_, err := r.db.Create(companyModel)
	if err != nil {
		return nil, err
	}

	return companyModel.ToDomain(), nil
}
