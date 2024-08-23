package profile

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateProfileUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	profileRepoMock := mock.NewMockIProfileRepository(ctrl)
	createProfileUseCase := NewCreateProfileUseCase(profileRepoMock)

	profileToCreate := factories.NewProfileFactory("gerente")
	createdProfile := factories.NewProfileFactory("gerente")

	t.Run("Success", func(t *testing.T) {
		profileRepoMock.EXPECT().Create(profileToCreate).Return(createdProfile, nil)

		profile, err := createProfileUseCase.Execute(profileToCreate)

		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, createdProfile, profile)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("Erro ao criar perfil")

		profileRepoMock.EXPECT().Create(profileToCreate).Return(nil, expectedError)

		profile, err := createProfileUseCase.Execute(profileToCreate)

		assert.Error(t, err)
		assert.Nil(t, profile)
		assert.Equal(t, expectedError, err)
	})
}
