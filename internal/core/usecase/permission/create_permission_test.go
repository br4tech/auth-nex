package permission

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreatePermissionUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepoMock := mock.NewMockIPermissionRepository(ctrl)
	createPermissionUseCase := NewCreatePermissionUseCase(permissionRepoMock)

	permissionToCreate := factories.NewPermissionFactory("update")
	createdPermission := factories.NewPermissionFactory("update")

	t.Run("Success", func(t *testing.T) {
		permissionRepoMock.EXPECT().Create(permissionToCreate).Return(createdPermission, nil)

		permission, err := createPermissionUseCase.Execute(permissionToCreate)

		assert.NoError(t, err)
		assert.NotNil(t, permission)
		assert.Equal(t, createdPermission, permission)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar permissão")

		permissionRepoMock.EXPECT().Create(permissionToCreate).Return(nil, expectedError)

		permission, err := createPermissionUseCase.Execute(permissionToCreate)

		assert.Error(t, err)
		assert.Nil(t, permission)
		assert.Equal(t, expectedError, err)
	})
}
