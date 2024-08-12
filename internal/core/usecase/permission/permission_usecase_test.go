package permission

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPermissionUseCase_CreateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	permissionRepoMock := mock.NewMockIProfileRepository(ctrl)
	permissionUseCase := NewPermissionUseCase(permissionRepoMock)

	permissionDomain := &domain.Profile{
		Name: "admin",
	}

	t.Run("sucess", func(t *testing.T) {
		permissionRepoMock.EXPECT().Create(gomock.Any()).Return(permissionDomain, nil)

		createdPermission, _ := permissionUseCase.Create("admin")

		assert.Equal(t, permissionDomain.Name, createdPermission.Name)
	})

	t.Run("error", func(t *testing.T) {
		permissionInvalid := errors.New("Name invalid")

		permissionRepoMock.EXPECT().Create(gomock.Any()).Return(nil, permissionInvalid)

		createdPermission, err := permissionUseCase.Create("admin")

		assert.Error(t, err)
		assert.Nil(t, createdPermission)
	})

}
