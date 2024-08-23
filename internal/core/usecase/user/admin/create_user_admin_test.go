package admin

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserAdminUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	createUserAdminUseCase := NewCreateUsereAdminUseCase(userRepoMock)

	userToCreate := factories.NewUserFactory("admin")
	createdUser := factories.NewUserFactory("admin")

	t.Run("Success", func(t *testing.T) {
		createdUser.Role = "admin"

		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createUserAdminUseCase.Execute(userToCreate)

		assert.NoError(t, err)
		assert.Equal(t, "admin", userToCreate.Role) // Verifica se o papel foi alterado para "admin"
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar usuário admin")

		userRepoMock.EXPECT().Create(userToCreate).Return(nil, expectedError)

		err := createUserAdminUseCase.Execute(userToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
