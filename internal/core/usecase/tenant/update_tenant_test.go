package tenant

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateTenantUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tenantRepoMock := mock.NewMockITenantRepository(ctrl)
	updateTenantUseCase := NewUpdateTenantUseCase(tenantRepoMock)

	tenantToUpdate := factories.NewTenantFactory()
	updatedTenant := factories.NewTenantFactory()

	t.Run("Success", func(t *testing.T) {
		updatedTenant.Name = "Nome Atualizado"

		tenantRepoMock.EXPECT().Update(tenantToUpdate).Return(updatedTenant, nil)

		tenant, err := updateTenantUseCase.Execute(tenantToUpdate)

		assert.NoError(t, err)
		assert.NotNil(t, tenant)
		assert.Equal(t, updatedTenant, tenant)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao atualizar tenant")

		tenantRepoMock.EXPECT().Update(tenantToUpdate).Return(nil, expectedError)

		tenant, err := updateTenantUseCase.Execute(tenantToUpdate)

		assert.Error(t, err)
		assert.Nil(t, tenant)
		assert.Equal(t, expectedError, err)
	})
}
