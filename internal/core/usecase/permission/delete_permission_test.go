package permission

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeletePermissionUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepoMock := mock.NewMockIPermissionRepository(ctrl)
	deletePermissionUseCase := NewDeletePermissionUseCase(permissionRepoMock)

	t.Run("Success", func(t *testing.T) {
		permissionId := 1
		permissionRepoMock.EXPECT().Delete(permissionId).Return(nil)

		err := deletePermissionUseCase.Execute(permissionId)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		permissionId := 1
		expectedError := errors.New("Erro ao excluir permissão")

		permissionRepoMock.EXPECT().Delete(permissionId).Return(expectedError)

		err := deletePermissionUseCase.Execute(permissionId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
