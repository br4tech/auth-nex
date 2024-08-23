package admin

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteUserAdminUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	deleteUserAdminUseCase := NewDeleteUserAdminUsecase(userRepoMock)

	t.Run("Success", func(t *testing.T) {
		userId := 1

		userRepoMock.EXPECT().Delete(userId).Return(nil)

		err := deleteUserAdminUseCase.Execute(userId)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		userId := 1
		expectedError := errors.New("Erro ao excluir usuário admin")

		userRepoMock.EXPECT().Delete(userId).Return(expectedError)

		err := deleteUserAdminUseCase.Execute(userId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
