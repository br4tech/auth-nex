package tenant

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTenantUseCase_CreateTenantWithCompanyAndAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tenantRepoMock := mock.NewMockITenantRepository(ctrl)
	companyUseCaseMock := mock.NewMockICompanyUseCase(ctrl)
	userUseCaseMock := mock.NewMockIUserUseCase(ctrl)

	cpfPtr := "123.456.789-00"
	profileIdPtr := 1

	tenantUseCase := NewTenantUseCase(tenantRepoMock, companyUseCaseMock, userUseCaseMock)

	tenantDTO := &dto.TenantDTO{
		Name: "app-0000",
		Company: dto.CompanyDTO{
			LegalName:         "Empresa Legal Ltda.",
			TradeName:         "Empresa Legal",
			Document:          "12345678901234",
			StateRegistration: "123456789",
			Address: dto.AddressDTO{
				Street:     "Rua Principal",
				Number:     "123",
				Complement: "Apto 45",
				District:   "Centro",
				City:       "São Paulo",
				State:      "SP",
				ZipCode:    "01234567",
			},
			Type: "LTDA",
		},
		User: dto.UserSystemDTO{
			Name:     "João Silva",
			Email:    "joao@email.com",
			CPF:      "12345678901",
			Password: "senha_forte",
			TenantId: 1,
			Roles:    "system",
		},
	}

	tenantDomain := &domain.Tenant{
		Name: "app-0000",
	}

	companyDomain := &domain.Company{
		LegalName:         "Empresa Legal Ltda.",
		TradeName:         "Empresa Legal",
		Document:          "12345678901234",
		StateRegistration: "123456789",
		Address: domain.Address{
			Street:     "Rua Principal",
			Number:     "123",
			Complement: "Apto 45",
			District:   "Centro",
			City:       "São Paulo",
			State:      "SP",
			ZipCode:    "01234567",
		},
		Type: "LTDA",
	}

	userDomain := &domain.User{
		Name:      "João Silva",
		Email:     "joao@email.com",
		CPF:       cpfPtr,
		Password:  "senha_forte",
		ProfileId: profileIdPtr,
		TenantId:  1,
	}

	t.Run("success", func(t *testing.T) {
		companyUseCaseMock.EXPECT().Create(gomock.Any()).Return(companyDomain, nil)
		userUseCaseMock.EXPECT().Create(gomock.Any()).Return(userDomain, nil)
		tenantRepoMock.EXPECT().Create(gomock.Any()).Return(tenantDomain, nil)

		createdTenant, _ := tenantUseCase.CreateTenantWithCompanyAndAdmin(tenantDTO)

		assert.Equal(t, tenantDTO.Name, createdTenant.Name)
		assert.Equal(t, tenantDTO.Company.LegalName, createdTenant.Companies[0].LegalName)
		assert.Equal(t, tenantDTO.Company.Document, createdTenant.Companies[0].Document)
		assert.Equal(t, tenantDTO.User.Name, createdTenant.Users[0].Name)
		assert.Equal(t, tenantDTO.User.Email, createdTenant.Users[0].Email)
	})

	t.Run("Error", func(t *testing.T) {

		invalidTenantDTO := &dto.TenantDTO{
			Name:    "app-0001",
			Company: dto.CompanyDTO{},
		}

		invalidTenant := errors.New("Invalid tenant")

		tenantRepoMock.EXPECT().Create(gomock.Any()).Return(nil, invalidTenant)

		createdTenant, err := tenantUseCase.CreateTenantWithCompanyAndAdmin(invalidTenantDTO)

		assert.Error(t, err)
		assert.Nil(t, createdTenant)
	})
}
