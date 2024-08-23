package tenant

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTenantByNameUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tenantRepoMock := mock.NewMockITenantRepository(ctrl)
	getTenantByNameUseCase := NewGetTenantByNameUseCase(tenantRepoMock)

	expectedTenant := factories.NewTenantFactory()

	t.Run("Success", func(t *testing.T) {
		tenantName := "Tenant Teste"

		tenantRepoMock.EXPECT().FindByName(tenantName).Return(expectedTenant, nil)

		tenant, err := getTenantByNameUseCase.Execute(tenantName)

		assert.NoError(t, err)
		assert.NotNil(t, tenant)
		assert.Equal(t, expectedTenant, tenant)
	})

	t.Run("Error - Tenant not found", func(t *testing.T) {
		tenantName := "Tenant Inexistente"
		expectedError := errors.New("Tenant não encontrado")

		tenantRepoMock.EXPECT().FindByName(tenantName).Return(nil, expectedError)

		tenant, err := getTenantByNameUseCase.Execute(tenantName)

		assert.Error(t, err)
		assert.Nil(t, tenant)
		assert.Equal(t, expectedError, err)
	})
}
