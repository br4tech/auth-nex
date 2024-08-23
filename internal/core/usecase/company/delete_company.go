package company

import "github.com/br4tech/auth-nex/internal/core/port"

type DeleteCompanyUsecase struct {
	companyRepo port.ICompanyRepository
}

func NewDeleteCompanyUseCase(companyRepo port.ICompanyRepository) *DeleteCompanyUsecase {
	return &DeleteCompanyUsecase{
		companyRepo: companyRepo,
	}
}

func (uc *DeleteCompanyUsecase) Execute(id int) error {
	return uc.companyRepo.Delete(id)
}
