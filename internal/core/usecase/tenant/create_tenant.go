package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
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
	tenantId, err := uc.createTenant(tenant)
	if err != nil {
		return err
	}

	tenant.Companies[0].TenantId = tenantId
	err = uc.createCompany(&tenant.Companies[0])
	if err != nil {
		return err
	}

	tenant.Users[0].TenantId = tenantId
	err = uc.createUserWithAdmin(&tenant.Users[0])
	if err != nil {
		return err
	}

	return nil
}

func (uc *CreateTenantUseCase) createTenant(tenant *domain.Tenant) (int, error) {

	tenant, err := uc.tenantRepository.Create(tenant)
	if err != nil {
		return 0, err
	}

	return tenant.Id, nil
}

func (uc *CreateTenantUseCase) createCompany(company *domain.Company) error {
	company.ParentCompanyId = 0 //seta empresa como matriz

	_, err := uc.companyRepository.Create(company)
	if err != nil {
		return err
	}

	return nil
}

func (uc *CreateTenantUseCase) createUserWithAdmin(user *domain.User) error {
	user.Role = "admin"

	_, err := uc.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
