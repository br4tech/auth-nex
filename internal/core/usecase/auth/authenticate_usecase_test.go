package auth

import (
	"testing"
	"time"

	"github.com/br4tech/auth-nex/internal/core/domain"
	"github.com/br4tech/auth-nex/internal/dto"
	"github.com/br4tech/auth-nex/internal/model"
	"github.com/br4tech/auth-nex/internal/test/mock"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticateUserUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock.NewMockIUserRepository(ctrl)
	AuthenticateUserUseCase := NewAuthenticateUserUseCase(userRepoMock)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

	userDomain := &domain.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	t.Run("Authenticate - Success", func(t *testing.T) {
		userRepoMock.EXPECT().FindByEmail(userDomain.Email).Return(userDomain, nil)

		token, err := AuthenticateUserUseCase.Authenticate(&dto.UserTokenDTO{
			Email:    userDomain.Email,
			Password: "123456",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := AuthenticateUserUseCase.ValidateAccessToken(*token)
		assert.NoError(t, err)
		assert.Equal(t, userDomain.Email, claims.Email)
	})

	t.Run("Authenticate - Invalid Password", func(t *testing.T) {
		userRepoMock.EXPECT().FindByEmail(userDomain.Email).Return(userDomain, nil)

		token, err := AuthenticateUserUseCase.Authenticate(&dto.UserTokenDTO{
			Email:    userDomain.Email,
			Password: "wrongpassword",
		})

		assert.Error(t, err)
		assert.Nil(t, token)
	})

	t.Run("CreateUser - Success", func(t *testing.T) {
		userRepoMock.EXPECT().Create(gomock.Any()).Return(nil, nil)

		createdUser, err := AuthenticateUserUseCase.Create(userDomain)

		assert.NoError(t, err)
		assert.Equal(t, userDomain, createdUser)
	})

	t.Run("ValidateAccessToken - Success", func(t *testing.T) {
		token, _ := generateAccessToken(userDomain.Email)

		claims, err := AuthenticateUserUseCase.ValidateAccessToken(token)

		assert.NoError(t, err)
		assert.Equal(t, userDomain.Email, claims.Email)
	})

	t.Run("ValidateAccessToken - Expired", func(t *testing.T) {
		claims := &model.Claims{
			Email: userDomain.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(), // Expired an hour ago
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(jwtKey)

		_, err := AuthenticateUserUseCase.ValidateAccessToken(tokenString)

		assert.Error(t, err)
	})
}
