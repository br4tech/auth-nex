package admin

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateUserAdminUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	updateUserAdminUseCase := NewUpdateUserAdminUsecase(userRepoMock)

	userToUpdate := factories.NewUserFactory("admin")
	updatedUser := factories.NewUserFactory("admin")

	t.Run("Success", func(t *testing.T) {
		updatedUser.Name = "Nome Atualizado"

		userRepoMock.EXPECT().Update(userToUpdate).Return(updatedUser, nil)

		user, err := updateUserAdminUseCase.Execute(userToUpdate)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, updatedUser, user)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao atualizar usuário admin")

		userRepoMock.EXPECT().Update(userToUpdate).Return(nil, expectedError)

		user, err := updateUserAdminUseCase.Execute(userToUpdate)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
