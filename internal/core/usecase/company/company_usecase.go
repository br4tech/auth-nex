package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/br4tech/auth-nex/pkg/validator"
	"github.com/jinzhu/copier"
)

type CompanyUseCase struct {
	companyRepository port.ICompanyRepository
}

func NewCompanyUseCase(companyRepository port.ICompanyRepository) port.ICompanyUseCase {
	return &CompanyUseCase{companyRepository: companyRepository}
}

func (uc *CompanyUseCase) FindCompanyById(id int) (*domain.Company, error) {
	company, err := uc.companyRepository.FindCompanyById(id)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (uc *CompanyUseCase) CreateCompany(company *domain.Company) (*domain.Company, error) {
	companyModel := &model.Company{}

	err := copier.Copy(companyModel, company)
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(companyModel); err != nil {
		return nil, err
	}

	createdCompany, err := uc.companyRepository.CreateCompany(companyModel.ToDomain())
	if err != nil {
		return nil, err
	}

	return createdCompany, nil
}
