package attendance

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserAttendanceByEmailUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	getUserAttendanceByEmailUseCase := NewGetUserAttendanceByEmailUsecase(userRepoMock)

	expectedUser := factories.NewUserFactory("Attendance")

	t.Run("Success", func(t *testing.T) {
		userEmail := "attendance@example.com"
		expectedUser.Email = userEmail
		expectedUser.Role = "attendance"

		userRepoMock.EXPECT().FindByEmail(userEmail).Return(expectedUser, nil)

		user, err := getUserAttendanceByEmailUseCase.Execute(userEmail)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		userEmail := "nonexistent@example.com"
		expectedError := errors.New("Usuário não encontrado")

		userRepoMock.EXPECT().FindByEmail(userEmail).Return(nil, expectedError)

		user, err := getUserAttendanceByEmailUseCase.Execute(userEmail)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
