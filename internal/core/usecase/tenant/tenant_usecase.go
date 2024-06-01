package tenant

import (
	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/br4tech/auth-nex/pkg/validator"

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

	createdTenant, err := uc.tenantRepository.CreateTenant(tenantModel.ToDomain())
	if err != nil {
		return nil, err
	}

	errCompany := copier.Copy(companyModel, tenant.Company)
	if errCompany != nil {
		return nil, errCompany
	}

	companyModel.TenantID = createdTenant.Id

	if err := validator.ValidateStruct(companyModel); err != nil {
		return nil, err
	}

	createdCompany, err := uc.companyUseCase.CreateCompany(companyModel.ToDomain())
	if err != nil {
		return nil, err
	}

	errUser := copier.Copy(userModel, tenant.User)
	if errUser != nil {
		return nil, errUser
	}
	companyModel.ID = uint(createdCompany.Id)
	userModel.TenantId = createdTenant.Id
	userModel.ProfileId = 1
	userModel.Companies = append(userModel.Companies, *companyModel)

	_, err = uc.userUsecase.CreateUser(userModel.ToDomain())

	if err != nil {
		return nil, err
	}

	tenantModel.Companies = append(tenantModel.Companies, *companyModel)
	tenantModel.Users = append(tenantModel.Users, *userModel)

	return tenantModel.ToDomain(), nil
}
