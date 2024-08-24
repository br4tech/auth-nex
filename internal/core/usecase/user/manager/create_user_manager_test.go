package manager

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserManagerUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	createUserManagerUseCase := NewCreateUserManagerUseCase(userRepoMock)

	userToCreate := factories.NewUserFactory("manager")
	createdUser := factories.NewUserFactory("manager")

	t.Run("Success", func(t *testing.T) {
		userRepoMock.EXPECT().Create(gomock.Any()).Return(createdUser, nil)

		err := createUserManagerUseCase.Execute(userToCreate)

		assert.NoError(t, err)
		assert.Equal(t, "manager", userToCreate.Role)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar usuário manager")

		userRepoMock.EXPECT().Create(userToCreate).Return(nil, expectedError)

		err := createUserManagerUseCase.Execute(userToCreate)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
