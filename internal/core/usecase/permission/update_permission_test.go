package permission

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdatePermissionUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepoMock := mock.NewMockIPermissionRepository(ctrl)
	updatePermissionUseCase := NewUpdatePermissionUseCase(permissionRepoMock)

	permissionToUpdate := factories.NewPermissionFactory("created")

	t.Run("Success", func(t *testing.T) {
		updatedPermission := &domain.Permission{
			Id:   permissionToUpdate.Id,
			Name: "updated_permission_name",
		}

		permissionRepoMock.EXPECT().Update(permissionToUpdate).Return(updatedPermission, nil)

		permission, err := updatePermissionUseCase.Execute(permissionToUpdate)

		assert.NoError(t, err)
		assert.NotNil(t, permission)
		assert.Equal(t, updatedPermission, permission)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao atualizar permissão")

		permissionRepoMock.EXPECT().Update(permissionToUpdate).Return(nil, expectedError)

		permission, err := updatePermissionUseCase.Execute(permissionToUpdate)

		assert.Error(t, err)
		assert.Nil(t, permission)
		assert.Equal(t, expectedError, err)
	})
}
