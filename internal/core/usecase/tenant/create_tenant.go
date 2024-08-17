package tenant

import (
	"errors"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/model"
)

type CreateTenantUseCase struct {
	tenantRepository  port.ITenantRepository
	companyRepository port.ICompanyRepository
	userRepository    port.IUserRepository
}

func NewCreateTenantUseCase(tenantRepository port.ITenantRepository,
	companyRepository port.ICompanyRepository, userRepository port.IUserRepository,
) *CreateTenantUseCase {
	return &CreateTenantUseCase{
		tenantRepository:  tenantRepository,
		userRepository:    userRepository,
		companyRepository: companyRepository,
	}
}

func (uc *CreateTenantUseCase) Execute(tenant *domain.Tenant) error {
	tenantModel := new(model.Tenant)
	tenantModel.FromDomain(tenant)

	return uc.tenantRepository.Create(tenantModel)
}

func (uc *CreateTenantUseCase) createTenant(tenant *domain.Tenant) error {
	return nil
}

func (uc *CreateTenantUseCase) createCompany(tenantId int, company *domain.Company) error {
	if tenantId == 0 {
		return errors.New("Tenant Id invalid")
	}

	company.TenantId = tenantId
	companyModel := new(model)

	return uc.companyRepository.Create(company)
}

func (uc *CreateTenantUseCase) createUserWithAdmin(companyId int, user *domain.User) error {
	return nil
}
