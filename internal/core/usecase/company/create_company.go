package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateCompanyUseCase struct {
	companyRepo port.ICompanyRepository
}

func NewCreateCompanyUseCase(companyRepo port.ICompanyRepository) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{
		companyRepo: companyRepo,
	}
}

func (uc *CreateCompanyUseCase) Execute(company *domain.Company) (*domain.Company, error) {
	company, err := uc.companyRepo.Create(company)
	if err != nil {
		return nil, err
	}

	return company, nil
}
