package system

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserSystemUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	createUserSystemUseCase := NewCreateUserSytemUseCase(userRepoMock)

	userToCreate := factories.NewUserFactory("system_admin")
	createdUser := factories.NewUserFactory("system_admin")

	t.Run("Success", func(t *testing.T) {
		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createUserSystemUseCase.Execute(userToCreate)

		assert.NoError(t, err)
		assert.Equal(t, "system_admin", userToCreate.Role)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar usuário system")

		userRepoMock.EXPECT().Create(userToCreate).Return(nil, expectedError)

		err := createUserSystemUseCase.Execute(userToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
