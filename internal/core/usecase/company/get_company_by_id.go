package company

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type GetCompanyByIdUseCase struct {
	companyRepo port.ICompanyRepository
}

func NewGetCompanyByIdUsecase(companyRepo port.ICompanyRepository) *GetCompanyByIdUseCase {
	return &GetCompanyByIdUseCase{
		companyRepo: companyRepo,
	}
}

func (uc *GetCompanyByIdUseCase) Execute(id int) (*domain.Company, error) {
	comapany, err := uc.companyRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return comapany.ToDomain(), nil
}
