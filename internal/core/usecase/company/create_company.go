package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type CreateCompanyUseCase struct {
	companyRepo port.ICompanyRepository
}

func NewCreateCompanyUseCase(companyRepo port.ICompanyRepository) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{
		companyRepo: companyRepo,
	}
}

func (uc *CreateCompanyUseCase) Execute(company *domain.Company) error {
	companyModel := new(model.Company)
	companyModel.FromDomain(company)

	return uc.companyRepo.Create(companyModel)
}
