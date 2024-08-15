package company

import "github.com/br4tech/auth-nex/internal/core/port"

type GetCompanyByIdUseCase struct {
	companyRepo port.ICompanyRepository
}

func NewGetCompanyByIdUsecase(companyRepo port.ICompanyRepository) *GetCompanyByIdUseCase {
	return &GetCompanyByIdUseCase{
		companyRepo: companyRepo,
	}
}

func (uc *GetCompanyByIdUseCase) Execute(id int) error {
	return nil
}
