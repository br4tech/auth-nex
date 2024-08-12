package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/core/validator"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/jinzhu/copier"
)

type CompanyUseCase struct {
	companyRepository port.ICompanyRepository
}

func NewCompanyUseCase(companyRepository port.ICompanyRepository) port.ICompanyUseCase {
	return &CompanyUseCase{companyRepository: companyRepository}
}

func (uc *CompanyUseCase) FindById(id int) (*domain.Company, error) {
	company, err := uc.companyRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (uc *CompanyUseCase) Create(company *domain.Company) (*domain.Company, error) {
	companyModel := &model.Company{}

	err := copier.Copy(companyModel, company)
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(companyModel); err != nil {
		return nil, err
	}

	createdCompany, err := uc.companyRepository.Create(companyModel.ToDomain())
	if err != nil {
		return nil, err
	}

	return createdCompany, nil
}
