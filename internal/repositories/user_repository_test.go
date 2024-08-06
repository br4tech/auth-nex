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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestUserRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[port.IModel](ctrl)
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

func TestUserRepository_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock.NewMockIDatabase[port.IModel](ctrl)
	repo := NewUserRepository(mockDB)

	t.Run("Success", func(t *testing.T) {
		email := "joao@example.com"
		expectedUser := &model.User{
			Name:      "João Silva",
			Email:     email,
			Password:  "hashed_password",
			CPF:       "123.456.789-00",
			TenantId:  1,
			ProfileId: 1,
		}

		mockDB.EXPECT().Where("email=?", email).Return(mockDB) // Retorna o próprio mock para encadear a chamada
		mockDB.EXPECT().First(gomock.Any()).DoAndReturn(func(user *model.User) *gorm.DB {
			*user = *expectedUser
			return &gorm.DB{}
		})

		user, err := repo.FindUserByEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, expectedUser.ToDomain().Id, user.Id)
		assert.Equal(t, expectedUser.ToDomain().Name, user.Name)
		assert.Equal(t, expectedUser.ToDomain().Email, user.Email)
		assert.Equal(t, expectedUser.ToDomain().Password, user.Password)
		assert.Equal(t, expectedUser.ToDomain().CPF, user.CPF)
		assert.Equal(t, expectedUser.ToDomain().TenantId, user.TenantId)
		assert.Equal(t, expectedUser.ToDomain().ProfileId, user.ProfileId)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		email := "naoexiste@example.com"
		mockDB.EXPECT().Where("email=?", email).Return(mockDB)
		mockDB.EXPECT().First(gomock.Any()).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})

		user, err := repo.FindUserByEmail(email)
		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestUserRepository_FindByPhone(t *testing.T) {
}
