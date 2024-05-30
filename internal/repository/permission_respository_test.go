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

	role := &domain.Role{
		Name: "Test Role",
	}

	roleModel := new(model.Role)
	roleModel.FromDomain(role)

	t.Run("Success", func(t *testing.T) {
		mockGormDB := &gorm.DB{}

		mockDB.EXPECT().Create(roleModel).Return(mockGormDB, nil)

		result, err := repo.CreateRole(role)
		if err != nil {
			t.Fatalf("CreateRole returned an error: %v", err)
		}

		if result == nil {
			t.Fatal("CreateRole returned a nil result")
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(roleModel).Return(nil, expectedError)

		_, err := repo.CreateRole(role)

		if err == nil {
			t.Fatal("CreateRole did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateRole returned unexpected error: %v", err)
		}
	})
}
