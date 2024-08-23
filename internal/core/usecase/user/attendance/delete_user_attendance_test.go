package attendance

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteUserAttendanceUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	deleteUserAttendanceUseCase := NewDeleteUserAttendanceUsecase(userRepoMock)

	t.Run("Success", func(t *testing.T) {
		userId := 1

		userRepoMock.EXPECT().Delete(userId).Return(nil)

		err := deleteUserAttendanceUseCase.Execute(userId)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		userId := 1
		expectedError := errors.New("Erro ao excluir usuário attendance")

		userRepoMock.EXPECT().Delete(userId).Return(expectedError)

		err := deleteUserAttendanceUseCase.Execute(userId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
