package customer

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteUserCustomerUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	deleteUserCustomerUseCase := NewDeleteUserCustomerUsecase(userRepoMock)

	t.Run("Success", func(t *testing.T) {
		userId := 1

		userRepoMock.EXPECT().Delete(userId).Return(nil)

		err := deleteUserCustomerUseCase.Execute(userId)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		userId := 1
		expectedError := errors.New("Erro ao excluir usuário customer")

		userRepoMock.EXPECT().Delete(userId).Return(expectedError)

		err := deleteUserCustomerUseCase.Execute(userId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
