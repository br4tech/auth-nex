package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
)

type CreateTenantUseCase struct {
	tenantRepository  port.ITenantRepository
	companyRepository port.ICompanyRepository
	userRepository    port.IUserRepository
	profileRepository port.IProfileRepository
}

func NewCreateTenantUseCase(tenantRepository port.ITenantRepository,
	companyRepository port.ICompanyRepository, userRepository port.IUserRepository,
	profileRepository port.IProfileRepository,
) port.ICreateTenantUseCase {
	return &CreateTenantUseCase{
		tenantRepository:  tenantRepository,
		userRepository:    userRepository,
		companyRepository: companyRepository,
		profileRepository: profileRepository,
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

	defaultProfile, err := uc.createDefaultProfile(tenantId)
	if err != nil {
		return err
	}

	tenant.Users[0].TenantId = tenantId
	tenant.Users[0].Profiles = []domain.Profile{{Id: defaultProfile.Id}}
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

func (uc *CreateTenantUseCase) createDefaultProfile(tenantId int) (*domain.Profile, error) {
	defaultProfile := &domain.Profile{
		Name:     "Administrador",
		TenantId: tenantId,
		Permisions: []domain.Permission{
			{Name: "criar_usuario"},
			{Name: "gerenciar_perfis"},
			{Name: "criar_empresas"},
		},
	}

	return uc.profileRepository.Create(defaultProfile)
}
