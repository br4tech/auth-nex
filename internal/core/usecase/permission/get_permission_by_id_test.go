package permission

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetPermissionById_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepoMock := mock.NewMockIPermissionRepository(ctrl)
	getPermissionByIdUseCase := NewGetPermissionById(permissionRepoMock)

	expectedPermission := factories.NewPermissionFactory("created")

	t.Run("Success", func(t *testing.T) {
		permissionId := 1
		permissionRepoMock.EXPECT().FindById(permissionId).Return(expectedPermission, nil)

		permission, err := getPermissionByIdUseCase.Execute(permissionId)

		assert.NoError(t, err)
		assert.NotNil(t, permission)
		assert.Equal(t, expectedPermission, permission)
	})

	t.Run("Error - Permission not found", func(t *testing.T) {
		permissionId := 1
		expectedError := errors.New("Permissão não encontrada")

		permissionRepoMock.EXPECT().FindById(permissionId).Return(nil, expectedError)

		permission, err := getPermissionByIdUseCase.Execute(permissionId)

		assert.Error(t, err)
		assert.Nil(t, permission)
		assert.Equal(t, expectedError, err)
	})
}
