package repository

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompany(db *gorm.DB) port.ICompanyRepository {
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

	if err := r.db.Create(&companyModel).Error; err != nil {
		return nil, err
	}

	return nil, nil
}
