package permission

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPermissionUseCase_CreateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepoMock := mock.NewMockIPermissionRepository(ctrl)
	permissionUseCase := NewPermissionUseCase(permissionRepoMock)

	permissionDomain := &domain.Role{
		Name: "admin",
	}

	t.Run("sucess", func(t *testing.T) {
		permissionRepoMock.EXPECT().CreateRole(gomock.Any()).Return(permissionDomain, nil)

		createdPermission, _ := permissionUseCase.CreateRole("admin")

		assert.Equal(t, permissionDomain.Name, createdPermission.Name)
	})

	t.Run("error", func(t *testing.T) {
		permissionInvalid := errors.New("Name invalid")

		permissionRepoMock.EXPECT().CreateRole(gomock.Any()).Return(nil, permissionInvalid)

		createdPermission, err := permissionUseCase.CreateRole("admin")

		assert.Error(t, err)
		assert.Nil(t, createdPermission)
	})

}
