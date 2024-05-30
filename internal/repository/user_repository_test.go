package repository

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestUserRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase(ctrl)
	repo := NewUserRepository(mockDB)

	user := &domain.User{
		Name:      "João Silva",
		Email:     "joao@example.com",
		Password:  "senha123",
		CPF:       "123.456.789-00",
		TenantId:  1,
		ProfileId: 2,
	}

	userModel := new(model.User)
	userModel.FromDomain(user)

	t.Run("Success", func(t *testing.T) {
		mockGormDB := &gorm.DB{}

		mockDB.EXPECT().Create(gomock.Any()).DoAndReturn(
			func(value interface{}) (*gorm.DB, error) {
				userModel := value.(*model.User)

				err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte("senha123"))
				if err != nil {
					t.Fatalf("Password not hashed: %v", err)
				}

				return mockGormDB, nil
			},
		)

		result, err := repo.CreateUser(user)
		if err != nil {
			t.Fatalf("CreateUser returned an error: %v", err)
		}

		if result == nil {
			t.Fatal("CreateUser returned a nil result")
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedError := errors.New("database error")

		mockDB.EXPECT().Create(userModel).Return(nil, expectedError)

		_, err := repo.CreateUser(user)

		if err == nil {
			t.Fatal("CreateUser did not return an error")
		}

		if err != expectedError {
			t.Fatalf("CreateUser returned unexpected error: %v", err)
		}
	})
}
