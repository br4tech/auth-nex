package attendance

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserAttendanceUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	createUserAttendanceUseCase := NewCreateUserAttendanceUsecase(userRepoMock)

	userToCreate := factories.NewUserFactory("attendance")
	createdUser := factories.NewUserFactory("attendance")

	t.Run("Success", func(t *testing.T) {
		createdUser.Role = "attendance"

		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createUserAttendanceUseCase.Execute(userToCreate)

		assert.NoError(t, err)
		assert.Equal(t, "attendance", userToCreate.Role)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar usuário attendance")

		userRepoMock.EXPECT().Create(userToCreate).Return(nil, expectedError)

		err := createUserAttendanceUseCase.Execute(userToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
