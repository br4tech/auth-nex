package profile

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetProfileByIdUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	profileRepoMock := mock.NewMockIProfileRepository(ctrl)
	getProfileByIdUseCase := NewGetProfileByIdUseCase(profileRepoMock)

	profileId := 1
	expectedProfile := factories.NewProfileFactory("secretaria")

	t.Run("Success", func(t *testing.T) {

		profileRepoMock.EXPECT().FindById(profileId).Return(expectedProfile, nil)

		profile, err := getProfileByIdUseCase.Execute(profileId)

		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, expectedProfile, profile)
	})

	t.Run("Error - Profile not found", func(t *testing.T) {
		expectedError := errors.New("Perfil não encontrado")

		profileRepoMock.EXPECT().FindById(profileId).Return(nil, expectedError)

		profile, err := getProfileByIdUseCase.Execute(profileId)

		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, expectedError, err)
	})
}
