package repository

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestPermissionRepository_CreateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase(ctrl)
	repo := NewPermissionRepository(mockDB)

	profile := &domain.Profile{
		Name: "Admin",
	}

	profileModel := new(model.Profile)
	profileModel.FromDomain(profile)

	t.Run("Success", func(t *testing.T) {
		mockGormDB := &gorm.DB{}

		mockDB.EXPECT().Create(profileModel).Return(mockGormDB, nil)

		result, err := repo.CreateProfile(profile)
		if err != nil {
			t.Fatalf("CreateProfile returned an error: %v", err)
		}

		if result == nil {
			t.Fatal("CreateProfile returned a nil result")
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(profileModel).Return(nil, expectedError)

		_, err := repo.CreateProfile(profile)

		if err == nil {
			t.Fatal("CreateProfile did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateProfile returned unexpected error: %v", err)
		}
	})
}
