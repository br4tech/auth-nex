package tenant

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTenantUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tenantRepoMock := mock.NewMockITenantRepository(ctrl)
	companyRepoMock := mock.NewMockICompanyRepository(ctrl)
	userRepoMock := mock.NewMockIUserRepository(ctrl)

	createTenantUseCase := NewCreateTenantUseCase(tenantRepoMock, companyRepoMock, userRepoMock)

	tenantToCreate := factories.NewTenantFactory()
	createdTenant := factories.NewTenantFactory()
	createdCompany := factories.NewCompanyFactory()
	createdUser := factories.NewUserFactory("admin")
	createdTenant.Id = 1

	t.Run("Success", func(t *testing.T) {
		companyToCreate := &tenantToCreate.Companies[0]
		userToCreate := &tenantToCreate.Users[0]

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(createdCompany, nil)
		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.NoError(t, err)
		assert.Equal(t, createdTenant.Id, companyToCreate.TenantId)
		assert.Equal(t, createdTenant.Id, userToCreate.TenantId)
	})

	t.Run("Error - Create Tenant fails", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar tenant")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - Create Company fails", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar empresa")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - Create User fails", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar u suário")

		tenantRepoMock.EXPECT().Create(tenantToCreate).Return(createdTenant, nil)
		companyRepoMock.EXPECT().Create(gomock.Any()).Return(createdCompany, nil)
		userRepoMock.EXPECT().Create(gomock.Any()).Return(nil, expectedError)

		err := createTenantUseCase.Execute(tenantToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
