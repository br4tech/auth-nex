package repository

import (
	"context"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type CompanyRepository struct {
	db port.IDatabase[model.Company]
}

func NewCompanyRepository(db port.IDatabase[model.Company]) port.ICompanyRepository {
	return &CompanyRepository{db: db}
}

func (r CompanyRepository) FindCompanyById(id int) (*domain.Company, error) {
	company, err := r.db.FindOne(context.Background(), "id=?", id)
	if err != nil {
		return nil, err
	}

	return company.ToDomain(), nil
}

func (r CompanyRepository) CreateCompany(company *domain.Company) (*domain.Company, error) {
	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	_, err := r.db.Create(context.Background(), companyModel)
	if err != nil {
		return nil, err
	}

	return companyModel.ToDomain(), nil
}
