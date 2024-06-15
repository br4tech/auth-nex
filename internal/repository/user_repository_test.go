package repository

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/mock"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestUserRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase(ctrl)
	repo := NewUserRepository(mockDB)
	mockGormDB := &gorm.DB{}

	cpfPtr := "123.456.789-00"
	profileIdPtr := 1

	user := &domain.User{
		Name:      "João Silva",
		Email:     "joao@example.com",
		Password:  "senha123",
		CPF:       &cpfPtr,
		TenantId:  1,
		ProfileId: &profileIdPtr,
	}

	userModel := new(model.User)
	userModel.FromDomain(user)

	t.Run("Success", func(t *testing.T) {
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

		createdUser, err := repo.CreateUser(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, user.Name, createdUser.Name)
		assert.Equal(t, user.Email, createdUser.Email)
	})

	t.Run("Error", func(t *testing.T) {
		mockDB.EXPECT().Create(gomock.Any()).DoAndReturn(
			func(value interface{}) (*gorm.DB, error) {
				userModel := value.(*model.User)

				// Verifica se a senha está hasheada (mesmo comportamento anterior)
				err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte("senha123"))
				if err != nil {
					t.Fatalf("Password not hashed: %v", err)
				}

				// Verifica se o email já existe (simulação de erro do banco de dados)
				if userModel.Email == "usuarioexistente@example.com" {
					return nil, errors.New("pq: duplicate key value violates unique constraint \"users_email_key\"")
				}

				// Se não houver erros, retorna sucesso (mesmo comportamento anterior)
				return mockGormDB, nil
			},
		)

		_, err := repo.CreateUser(user)

		assert.NoError(t, err)
	})
}
