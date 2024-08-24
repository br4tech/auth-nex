package manager

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserManagerByEmailUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	getUserManagerByEmailUseCase := NewGetUserManagerByEmailUseCase(userRepoMock)

	expectedUser := factories.NewUserFactory("manager")

	t.Run("Success", func(t *testing.T) {
		userEmail := "manager@example.com"
		expectedUser.Email = userEmail

		filter := map[string]interface{}{"email": userEmail}

		userRepoMock.EXPECT().FindBy(filter).Return(expectedUser, nil)

		user, err := getUserManagerByEmailUseCase.Execute(userEmail)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		userEmail := "nonexistent@example.com"
		expectedError := errors.New("Usuário não encontrado")

		filter := map[string]interface{}{"email": userEmail}

		userRepoMock.EXPECT().FindBy(filter).Return(nil, expectedError)

		user, err := getUserManagerByEmailUseCase.Execute(userEmail)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
