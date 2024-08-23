package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateCompanyUseCase struct {
	companyRepo port.ICompanyRepository
}

func NewUpdateCompanyUseCase(companyRepo port.ICompanyRepository) *UpdateCompanyUseCase {
	return &UpdateCompanyUseCase{
		companyRepo: companyRepo,
	}
}

func (uc *UpdateCompanyUseCase) Execute(company *domain.Company) (*domain.Company, error) {

	company, err := uc.companyRepo.Update(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}
