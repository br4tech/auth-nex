package profile

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/mock" // Certifique-se que o caminho está correto
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteProfileUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	profileRepoMock := mock.NewMockIProfileRepository(ctrl) // Certifique-se que o mock está implementado corretamente
	deleteProfileUseCase := NewDeleteProfileUseCase(profileRepoMock)

	t.Run("Success", func(t *testing.T) {
		profileId := 1

		profileRepoMock.EXPECT().Delete(profileId).Return(nil)

		err := deleteProfileUseCase.Execute(profileId)

		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		profileId := 1
		expectedError := errors.New("Erro ao excluir perfil")

		profileRepoMock.EXPECT().Delete(profileId).Return(expectedError)

		err := deleteProfileUseCase.Execute(profileId)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)
	})
}
