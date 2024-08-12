package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/core/validator"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"

	"github.com/jinzhu/copier"
)

type TenantUseCase struct {
	tenantRepository port.ITenantRepository
	companyUseCase   port.ICompanyUseCase
	userUsecase      port.IUserUseCase
}

func NewTenantUseCase(
	tenantRepository port.ITenantRepository,
	companyUseCase port.ICompanyUseCase,
	userUsecase port.IUserUseCase,
) port.ITenantUseCase {
	return &TenantUseCase{
		tenantRepository: tenantRepository,
		companyUseCase:   companyUseCase,
		userUsecase:      userUsecase,
	}
}

func (uc *TenantUseCase) CreateTenantWithCompanyAndAdmin(tenant *dto.TenantDTO) (*domain.Tenant, error) {
	tenantModel := &model.Tenant{}
	companyModel := &model.Company{}
	userModel := &model.User{}

	errTenant := copier.Copy(tenantModel, tenant)
	if errTenant != nil {
		return nil, errTenant
	}

	if err := validator.ValidateStruct(tenantModel); err != nil {
		return nil, err
	}

	createdTenant, err := uc.tenantRepository.Create(tenantModel.ToDomain())
	if err != nil {
		return nil, err
	}

	errCompany := copier.Copy(companyModel, tenant.Company)
	if errCompany != nil {
		return nil, errCompany
	}

	companyModel.TenantId = createdTenant.Id

	if err := validator.ValidateStruct(companyModel); err != nil {
		return nil, err
	}

	createdCompany, err := uc.companyUseCase.Create(companyModel.ToDomain())
	if err != nil {
		return nil, err
	}

	errUser := copier.Copy(userModel, tenant.User)
	if errUser != nil {
		return nil, errUser
	}
	companyModel.Id = createdCompany.Id
	userModel.TenantId = createdTenant.Id
	userModel.ProfileId = 1
	// userModel.Companies = append(userModel.Companies, *companyModel)

	_, err = uc.userUsecase.Create(userModel.ToDomain())

	if err != nil {
		return nil, err
	}

	tenantModel.Companies = append(tenantModel.Companies, *companyModel)
	tenantModel.Users = append(tenantModel.Users, *userModel)

	return tenantModel.ToDomain(), nil
}
