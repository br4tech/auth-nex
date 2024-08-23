package tenant

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTenantByIdUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tenantRepoMock := mock.NewMockITenantRepository(ctrl)
	getTenantByIdUseCase := NewGetTenantByIdUseCase(tenantRepoMock)

	expectedTenant := factories.NewTenantFactory()
	tenantId := 1

	t.Run("Success", func(t *testing.T) {
		tenantRepoMock.EXPECT().FindById(tenantId).Return(expectedTenant, nil)

		tenant, err := getTenantByIdUseCase.Execute(tenantId)

		assert.NoError(t, err)
		assert.NotNil(t, tenant)
		assert.Equal(t, expectedTenant, tenant)
	})

	t.Run("Error - Tenant not found", func(t *testing.T) {
		expectedError := errors.New("Tenant não encontrado")

		tenantRepoMock.EXPECT().FindById(tenantId).Return(nil, expectedError)

		tenant, err := getTenantByIdUseCase.Execute(tenantId)

		assert.Error(t, err)
		assert.Nil(t, tenant)
		assert.Equal(t, expectedError, err)
	})
}
