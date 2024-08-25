package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/br4tech/auth-nex/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestValidateTokenUseCase_Execute(t *testing.T) {
	jwtKey := []byte("chave_secreta_do_jwt")
	validateTokenUseCase := NewValidateTokenUseCase(jwtKey)

	t.Run("Success - Valid token", func(t *testing.T) {
		userId := 1
		claims := &model.Claims{
			Id: userId,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		assert.NoError(t, err)

		parsedClaims, err := validateTokenUseCase.Execute(tokenString)

		assert.NoError(t, err)
		assert.NotNil(t, parsedClaims)
		assert.Equal(t, userId, parsedClaims.Id)
	})

	t.Run("Error - Invalid token", func(t *testing.T) {
		invalidTokenString := "token_invalido"

		parsedClaims, err := validateTokenUseCase.Execute(invalidTokenString)

		assert.Error(t, err)
		assert.Nil(t, parsedClaims)
		assert.Equal(t, "Token inválido", err.Error())
	})

	t.Run("Error - Expired token", func(t *testing.T) {
		userId := 1
		claims := &model.Claims{
			Id: userId,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(), // Token expirado
				IssuedAt:  time.Now().Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		assert.NoError(t, err)

		parsedClaims, err := validateTokenUseCase.Execute(tokenString)

		assert.Error(t, err)
		assert.Nil(t, parsedClaims)
		assert.True(t, errors.Is(err, jwt.ErrTokenExpired))
	})
}
