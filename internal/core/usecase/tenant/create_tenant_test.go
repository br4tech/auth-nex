package tenant

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTenantUseCase_Execute(t *testing.T) {
	createdTenant := factories.NewTenantFactory()
	createdTenant.Id = 1
	createdCompany := factories.NewCompanyFactory()
	createdUser := factories.NewUserFactory("admin")

	defaultProfile := &domain.Profile{
		Name:     "Administrador",
		TenantId: createdTenant.Id,
		Permisions: []domain.Permission{
			{Name: "criar_usuario"},
			{Name: "gerenciar_perfis"},
			{Name: "criar_empresas"},
		},
	}

	t.Run("Success", func(t *testing.T) {
		ctrl := gomock.NewController(t) // Novo controlador para este cenário
		defer ctrl.Finish()

		tenantRepoMock := mock.NewMockITenantRepository(ctrl)
		companyRepoMock := mock.NewMockICompanyRepository(ctrl)
		userRepoMock := mock.NewMockIUserRepository(ctrl)
		profileRepoMock := mock.NewMockIProfileRepository(ctrl)

		createTenantUseCase := NewCreateTenantUseCase(
			tenantRepoMock, companyRepoMock,
			userRepoMock, profileRepoMock,
		)

		tenantToCreate := factories.NewTenantFactory()

		companyToCreate := &tenantToCreate.Companies[0]
		userToCreate := &tenantToCreate.Users[0]

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(createdCompany, nil)
		profileRepoMock.EXPECT().Create(gomock.Any()).Return(defaultProfile, nil)
		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.NoError(t, err)
		assert.Equal(t, createdTenant.Id, companyToCreate.TenantId)
		assert.Equal(t, createdTenant.Id, userToCreate.TenantId)
		assert.Equal(t, 1, len(userToCreate.Profiles))
		assert.Equal(t, defaultProfile.Id, userToCreate.Profiles[0].Id)
	})

	t.Run("Error - Create Tenant fails", func(t *testing.T) {
		ctrl := gomock.NewController(t) // Novo controlador para este cenário
		defer ctrl.Finish()

		tenantRepoMock := mock.NewMockITenantRepository(ctrl)

		createTenantUseCase := NewCreateTenantUseCase(
			tenantRepoMock, nil, nil, nil,
		)

		tenantToCreate := factories.NewTenantFactory()
		expectedError := errors.New("Erro ao criar tenant")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - Create Company fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tenantRepoMock := mock.NewMockITenantRepository(ctrl)
		companyRepoMock := mock.NewMockICompanyRepository(ctrl)

		createTenantUseCase := NewCreateTenantUseCase(
			tenantRepoMock, companyRepoMock, nil, nil,
		)

		tenantToCreate := factories.NewTenantFactory()
		expectedError := errors.New("Erro ao criar empresa")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - Create Default Profile fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tenantRepoMock := mock.NewMockITenantRepository(ctrl)
		companyRepoMock := mock.NewMockICompanyRepository(ctrl)
		profileRepoMock := mock.NewMockIProfileRepository(ctrl)

		createTenantUseCase := NewCreateTenantUseCase(
			tenantRepoMock, companyRepoMock, nil, profileRepoMock,
		)

		tenantToCreate := factories.NewTenantFactory()
		expectedError := errors.New("Erro ao criar perfil padrão")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(createdCompany, nil)
		profileRepoMock.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - Create User fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tenantRepoMock := mock.NewMockITenantRepository(ctrl)
		companyRepoMock := mock.NewMockICompanyRepository(ctrl)
		userRepoMock := mock.NewMockIUserRepository(ctrl)
		profileRepoMock := mock.NewMockIProfileRepository(ctrl)

		createTenantUseCase := NewCreateTenantUseCase(
			tenantRepoMock, companyRepoMock, userRepoMock, profileRepoMock,
		)

		tenantToCreate := factories.NewTenantFactory()
		expectedError := errors.New("Erro ao criar usuário")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(createdCompany, nil)
		profileRepoMock.EXPECT().Create(gomock.Any()).Return(defaultProfile, nil)
		userRepoMock.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
