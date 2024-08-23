package admin

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserAdminByEmailUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	getUserAdminByEmailUseCase := NewGetUserAdminByEmailUseCase(userRepoMock)

	expectedUser := factories.NewUserFactory("admin")

	t.Run("Success", func(t *testing.T) {
		userEmail := "admin@example.com"
		expectedUser.Email = userEmail // Garantindo que o email do usuário seja o mesmo que estamos buscando

		// Configura o mock para retornar o usuário esperado
		userRepoMock.EXPECT().FindByEmail(userEmail).Return(expectedUser, nil)

		user, err := getUserAdminByEmailUseCase.Execute(userEmail)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("Error - User not found", func(t *testing.T) {
		userEmail := "nonexistent@example.com"
		expectedError := errors.New("Usuário não encontrado")

		userRepoMock.EXPECT().FindByEmail(userEmail).Return(nil, expectedError)

		user, err := getUserAdminByEmailUseCase.Execute(userEmail)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, expectedError, err)
	})
}
