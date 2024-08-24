package profile

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateProfileUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	profileRepoMock := mock.NewMockIProfileRepository(ctrl)
	updateProfileUseCase := NewUpdateProfile(profileRepoMock)

	profileToUpdate := factories.NewProfileFactory("gerente")
	updatedProfile := factories.NewProfileFactory("atendendte")

	t.Run("Success", func(t *testing.T) {
		updatedProfile.Name = "Nome Atualizado"

		profileRepoMock.EXPECT().Update(profileToUpdate).Return(updatedProfile, nil)

		profile, err := updateProfileUseCase.Execute(profileToUpdate)

		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, updatedProfile, profile)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao atualizar perfil")

		// Configura o mock para retornar um erro
		profileRepoMock.EXPECT().Update(profileToUpdate).Return(nil, expectedError)

		profile, err := updateProfileUseCase.Execute(profileToUpdate)

		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, expectedError, err)
	})
}
