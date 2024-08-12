package repositories

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/core/port"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestPermissionRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[port.IModel](ctrl)
	repo := NewProfileRepository(mockDB)

	profile := &domain.Profile{
		Id:   1,
		Name: "Admin",
	}

	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	t.Run("Success", func(t *testing.T) {
		mockDB.EXPECT().Create(profileModel).Return(profile.Id, nil)

		result, err := repo.Create(profile)
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(profileModel).Return(0, expectedError)

		_, err := repo.Create(profile)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}

func TestPermissionRepository_FindByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[model.Profile](ctrl)
	repo := NewProfileRepository(mockDB)

	t.Run("Sucess", func(t *testing.T) {
		name := "Admin"

		expectedProfiles := &model.Profile{
			Name: name,
		}

		mockDB.EXPECT().FindBy("name= ?", name).Return(expectedProfiles, nil)

		profile, err := repo.FindByName(name)
		assert.NoError(t, err)
		assert.NotNil(t, profile)
		assert.Equal(t, expectedProfiles.Name, profile.Name)
	})

	t.Run("ProfileNotFound", func(t *testing.T) {
		name := "TesteA"
		expectedError := errors.New("ProfileNotFound")

		mockDB.EXPECT().FindBy("name= ?", name).Return(nil, expectedError)
		profile, err := repo.FindByName(name)
		assert.Error(t, err)
		assert.Nil(t, profile)
	})
}
