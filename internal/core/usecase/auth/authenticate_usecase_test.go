package auth

import (
	"errors"
	"testing"

	"github.com/br4tech/auth-nex/internal/core/validator"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/test/factories"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthenticateUserUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	generateTokenUseCaseMock := mock.NewMockIGenerateTokenUseCase(ctrl)

	authenticateUserUseCase := &AuthenticateUserUseCase{
		userRepository:       userRepoMock,
		generateTokenUseCase: generateTokenUseCaseMock,
	}

	t.Run("Success - Email and Password", func(t *testing.T) {
		userReq := &dto.UserTokenDTO{
			Email:    "test@example.com",
			Password: "senha123",
		}
		expectedUser := factories.NewUserFactory("admin")
		expectedUser.Email = userReq.Email
		expectedUser.Password = validator.HashPassword(userReq.Password)
		expectedToken := "token_valido"

		userRepoMock.EXPECT().FindByEmail(userReq.Email).Return(expectedUser, nil)
		generateTokenUseCaseMock.EXPECT().Execute(expectedUser.Id).Return(expectedToken, nil)

		token, err := authenticateUserUseCase.Execute(userReq)

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, expectedToken, *token)
	})

	t.Run("Success - Phone", func(t *testing.T) {
		userReq := &dto.UserTokenDTO{
			Phone: "+5511987654321",
		}
		expectedUser := factories.NewUserFactory("customer")
		expectedUser.Phone = userReq.Phone
		expectedToken := "token_valido"

		userRepoMock.EXPECT().FindByPhone(userReq.Phone).Return(expectedUser, nil) // Adaptado para o método FindBy
		generateTokenUseCaseMock.EXPECT().Execute(expectedUser.Id).Return(expectedToken, nil)

		token, err := authenticateUserUseCase.Execute(userReq)

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.Equal(t, expectedToken, *token)
	})

	t.Run("Error - Invalid email/password", func(t *testing.T) {
		userReq := &dto.UserTokenDTO{
			Email:    "test@example.com",
			Password: "senha_errada",
		}
		expectedUser := factories.NewUserFactory("admin")
		expectedUser.Email = userReq.Email
		expectedUser.Password = validator.HashPassword("senha123")

		userRepoMock.EXPECT().FindByEmail(userReq.Email).Return(expectedUser, nil)

		token, err := authenticateUserUseCase.Execute(userReq)

		assert.Error(t, err)
		assert.Nil(t, token)
		assert.Equal(t, "Incorrect password", err.Error())
	})

	t.Run("Error - User not found", func(t *testing.T) {
		userReq := &dto.UserTokenDTO{
			Email:    "nonexistent@example.com",
			Password: "senha123",
		}
		expectedError := errors.New("Usuário não encontrado")

		userRepoMock.EXPECT().FindByEmail(userReq.Email).Return(nil, expectedError)

		token, err := authenticateUserUseCase.Execute(userReq)

		assert.Error(t, err)
		assert.Nil(t, token)
		assert.Equal(t, expectedError, err)
	})

	t.Run("Error - Generate token fails", func(t *testing.T) {
		userReq := &dto.UserTokenDTO{
			Email:    "test@example.com",
			Password: "senha123",
		}
		expectedUser := factories.NewUserFactory("admin")
		expectedUser.Email = userReq.Email
		expectedUser.Password = validator.HashPassword(userReq.Password)
		expectedError := errors.New("Erro ao gerar token")

		userRepoMock.EXPECT().FindByEmail(userReq.Email).Return(expectedUser, nil)
		generateTokenUseCaseMock.EXPECT().Execute(expectedUser.Id).Return("", expectedError)

		token, err := authenticateUserUseCase.Execute(userReq)

		assert.Error(t, err)
		assert.Nil(t, token)
		assert.Equal(t, expectedError, err)
	})
}
