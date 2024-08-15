package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type UpdateCompanyUseCase struct {
	companyRepo port.ICompanyRepository
}

func NewUpdateCompany(companyRepo port.ICompanyRepository) *UpdateCompanyUseCase {
	return &UpdateCompanyUseCase{
		companyRepo: companyRepo,
	}
}

func (uc *UpdateCompanyUseCase) Execute(company *domain.Company) error {
	return nil
}
