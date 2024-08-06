package repositories

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type CompanyRepository[T port.IModel] struct {
	db port.IDatabase[T]
}

func NewCompanyRepository[T port.IModel](db port.IDatabase[T]) port.ICompanyRepository {
	return &CompanyRepository[T]{db: db}
}

func (r *CompanyRepository[T]) FindById(id int) (*domain.Company, error) {
	var companyEntity T

	_, err := r.db.FindById(id)
	if err != nil {
		return nil, err
	}

	companyModel, ok := any(&companyEntity).(*model.Company)
	if !ok {
		return nil, errors.New("failed to convert entity to company model")
	}

	companyDomain := companyModel.ToDomain()

	return companyDomain, nil
}

func (r *CompanyRepository[T]) Create(company *domain.Company) (*domain.Company, error) {
	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	companyEntity, ok := any(companyModel).(T)
	if !ok {
		return nil, errors.New("entity type invalid to company")
	}

	companyId, err := r.db.Create(companyEntity)
	if err != nil {
		return nil, err
	}
	company.Id = companyId

	return company, nil
}
